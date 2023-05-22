package ssti

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)



func GetPluginsFolder() string {
	var folder string
	// if folder ./plugins exists, use it
	if _, err := os.Stat("./plugins"); err == nil {
		folder = "./plugins"
	} else {
		// if folder ./plugins does not exist, use /opt/gossti/plugins
		folder = "/opt/gossti/plugins"
		// if folder /opt/gossti/plugins does not exist, create it
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			log.Debugf("Creating folder %s", folder)
			err = os.MkdirAll(folder, 0755)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	}

	log.Debugf("Using plugins folder: %s", folder)

	return folder
}
func getLanguage(name string) (Language, error) {
	path := fmt.Sprintf("%s/%s.yml", GetPluginsFolder(), name)
	log.Tracef("Load language %s from plugin %s", name, path)

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	language := Language{}
	err = yaml.Unmarshal([]byte(file), &language)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Debugf("Sucessfully loaded plugin for language: %v", language.Name)
	
	return language, nil
}

func getAvailableLanguages() []string {
	log.Trace("Get available languages from plugins folder")

	files, err := os.ReadDir(GetPluginsFolder())
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	languages := []string{}
	for _, file := range files {
		languages = append(languages, strings.TrimSuffix(file.Name(), ".yml"))
	}

	return languages
}