package cmd

import (
	"fmt"
	"net/http"
	"strings"

	gossti "github.com/LeoFVO/gossti/pkg"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// newCmd represents the build command
var newCmd = &cobra.Command{
	Use:   "scan",
	Short: "Try detect an SSTI vulnerability",
	Long:  `Try detect an SSTI vulnerability`,
	Args: func (cmd *cobra.Command, args []string) error {
		// Check if the payload contains the placeholder
		injectablePayload := cmd.Flags().Lookup("payload").Value.String()
		if !strings.Contains(injectablePayload, "SSTI") {
			return fmt.Errorf("payload should contain \"SSTI\" as a placeholder")
		}

		// check if the language is supported
		l := cmd.Flags().Lookup("languages").Value.String()
		if l != "all" {
			languages := strings.Split(l, ",")
			for _, language := range languages {
				if !gossti.IsLanguageSupported(language) {
					return fmt.Errorf("language %s not supported", language)
				}
			}
		}

		// Check if the method is valid
		method := cmd.Flags().Lookup("method").Value.String()
		if method != "GET" && method != "POST" && method != "PATCH" && method != "PUT" && method != "DELETE" && method != "TRACE" {
			return fmt.Errorf("method should be one of the following: GET, POST, PATCH, PUT, DELETE, TRACE")
		}

		// Check if the URL start with http or https
		url := cmd.Flags().Lookup("url").Value.String()
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			return fmt.Errorf("URL should start with http:// or https://")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flags().Lookup("url").Value.String()
		injectablePayload := cmd.Flags().Lookup("payload").Value.String()
		method := cmd.Flags().Lookup("method").Value.String()
		userAgent := cmd.Flags().Lookup("user-agent").Value.String()

		log.Tracef("Setting up the request")
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			panic(err)
		}
		log.Tracef("Setting up the user-agent")
		req.Header.Set("User-Agent", userAgent)

		log.Tracef("Getting the name and value of the payload")
		name := strings.Split(injectablePayload, "=")[0]
		value := strings.Split(injectablePayload, "=")[1]
		
		log.Tracef("Getting the languages to test against")
		l := cmd.Flags().Lookup("languages").Value.String()
		var languages []string

		if l == "all" {
			languages = gossti.GetSupportedLanguages()
		} else {
			languages = strings.Split(l, ",")
		}
		log.Tracef("Languages to test against: %s", languages)

		for _, language := range languages {
			log.Tracef("Testing against %s", language)
			language, err := gossti.Factory(language)
			if err != nil {
				panic(err)
			}
			language.Detect(req, gossti.QueryString{ Name: name, Value: value })
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("url", "u", "", "The target IP or domain to scan")
	newCmd.MarkFlagRequired("url")

	newCmd.Flags().StringP("payload", "p", "", "Payload containing the injection point, should contain \"SSTI\" as a placeholder")
	newCmd.MarkFlagRequired("payload")

	newCmd.Flags().StringP("method", "X", "GET", "The HTTP method to use")

	newCmd.Flags().StringP("user-agent", "", "gossti 1.0.0", "Custom user-agent to use")

	newCmd.Flags().StringP("languages", "l", "all", fmt.Sprintf("comma-separated languages names, can be one of the following: all, %s", strings.Join(gossti.GetSupportedLanguages(), ", ")))

}
