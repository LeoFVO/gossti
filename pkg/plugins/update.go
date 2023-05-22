package plugins

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/LeoFVO/gossti/pkg/ssti"
	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
)

const (
	pluginsFolderEndpoint = "https://api.github.com/repos/leofvo/gossti/git/trees/main?recursive=1"
)

type (
	RuntimesResponse struct {
		Sha  string `json:"sha"`
		URL  string `json:"url"`
		Tree []struct {
			Path string `json:"path"`
			Mode string `json:"mode"`
			Type string `json:"type"`
			Size int    `json:"size"`
			Sha  string `json:"sha"`
			URL  string `json:"url"`
		} `json:"tree"`
		Truncated bool `json:"truncated"`
	}
)

func UpdatePlugins() error {
	log.Info("Updating plugins...")

log.Debugf("GET request on '%s'", pluginsFolderEndpoint)
	response, err := http.Get(pluginsFolderEndpoint)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var runtimes RuntimesResponse
	err = json.NewDecoder(response.Body).Decode(&runtimes)
	if err != nil {
		return err
	}

	for _, tree := range runtimes.Tree {
		if tree.Type == "blob" && strings.HasPrefix(tree.Path, "plugins/") && strings.HasSuffix(tree.Path, ".yml") {
			log.Debugf("Found plugin '%s'", strings.Split(tree.Path, "/")[1])
			cli := gentleman.New()
			cli.URL("https://raw.githubusercontent.com/leofvo/gossti/main/" + tree.Path)
			res, err := cli.Request().Send()
			if err != nil {
				log.Errorf("Request error: %v\n", err)
				return err
			}
			if !res.Ok {
				log.Errorf("Invalid server response: %d\n", res.StatusCode)
				return err
			}
			filename := strings.Split(tree.Path, "/")[1]
			log.Debugf("Writing file '%s'", filename)
			os.WriteFile(ssti.GetPluginsFolder() + "/" + filename, res.Bytes(), 0644)
		}
	}
	
	log.Warn("All plugins updated")
	return nil
}
