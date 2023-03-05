package gossti

import "fmt"

// Factory pattern to create a language
func Factory(languageType string) (ILanguage, error) {
    switch languageType {
    case "nodeJS":
        return newNode(), nil
    case "ruby":
        return newRuby(), nil
    case "python":
        return newPython(), nil
    default:
        return nil, fmt.Errorf("language %s not supported", languageType)
    }
}

func GetSupportedLanguages() []string {
    return []string{"nodeJS", "ruby", "python"}
}

func IsLanguageSupported(languageType string) bool {
    for _, language := range GetSupportedLanguages() {
        if language == languageType {
            return true
        }
    }
    return false
}


