package plugins

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
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

	pluginsUpdated := 0

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

			remoteVersion := res.String()[strings.Index(res.String(), "version:")+9 : strings.Index(res.String(), "name:")-1]
			localVersion := ssti.GetPluginsVersion(strings.Split(strings.Split(tree.Path, "/")[1], ".")[0])

			filename := strings.Split(tree.Path, "/")[1]
			
			if isVersionGreater(remoteVersion, localVersion){
				log.Warnf("Plugin '%s' is outdated, updating", filename)
				log.Debugf("Writing file '%s'", filename)
				err := os.WriteFile(ssti.GetPluginsFolder() + "/" + filename, res.Bytes(), 0644)
				if err != nil {
					log.Fatal(err)
				}
				pluginsUpdated++
				} else {
					log.Infof("Plugin '%s' is up to date", filename)
					continue
				}
		}
	}
	
	log.Warnf("Updated %d plugins", pluginsUpdated)

	return nil
}


func isVersionGreater(version1 string, version2 string) bool {
	// Splitting version strings into individual components
	v1Components := strings.Split(version1, ".")
	v2Components := strings.Split(version2, ".")

	// Iterating over the components to compare
	for i := 0; i < len(v1Components) && i < len(v2Components); i++ {
		// Parsing the component values as integers
		v1Value, _ := strconv.Atoi(v1Components[i])
		v2Value, _ := strconv.Atoi(v2Components[i])

		// Comparing the component values
		if v1Value > v2Value {
			return true
		} else if v1Value < v2Value {
			return false
		}
	}

	// If all the common components are equal, checking if any remaining components exist in the first version
	for i := len(v1Components); i < len(v2Components); i++ {
		v2Value, _ := strconv.Atoi(v2Components[i])
		if v2Value > 0 {
			return false
		}
	}

	return false
}