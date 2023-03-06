package gossti

import (
	"fmt"
	"strings"
)

// Factory pattern to create a language
func Factory(languageType string) (ILanguage, error) {
    switch strings.ToLower(languageType) {
    case "nodejs", "node":
        return newNode(), nil
    case "ruby":
        return newRuby(), nil
    case "python":
        return newPython(), nil
    case "java":
        return newJava(), nil
    case "php":
        return newPhp(), nil
    default:
        return nil, fmt.Errorf("language %s not supported", languageType)
    }
}

func GetSupportedLanguages() []string {
    return []string{"nodejs", "node", "ruby", "python", "java", "php"}
}

func IsLanguageSupported(languageType string) bool {
    for _, language := range GetSupportedLanguages() {
        if language == strings.ToLower(languageType) {
            return true
        }
    }
    return false
}


