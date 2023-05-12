package ssti

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)


var (
	PLUGINS_FOLDER = "./plugins"
)

func getLanguage(name string) (Language, error) {
	path := fmt.Sprintf("%s/%s.yml", PLUGINS_FOLDER, name)
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

	files, err := os.ReadDir(PLUGINS_FOLDER)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	languages := []string{}
	for _, file := range files {
		languages = append(languages, strings.TrimSuffix(file.Name(), ".yml"))
	}

	return languages
}