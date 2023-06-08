package detect

import (
	"net/url"

	"github.com/LeoFVO/gossti/pkg/ssti"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:     "detect -u <url>",
		Aliases: []string{"scan"},
		Short: "Detect SSTI vulnerabilities",
		Long:  `Detect SSTI vulnerabilities in a given URL`,
		Example: `gossti detect -u http://example.com/?param1=SSTI&param2=value2`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Verify that the URL is valid
			u, _ := cmd.Flags().GetString("url")
			_, err := url.ParseRequestURI(u)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")

			// Define options
			options := ssti.Options{}

			method, _ := cmd.Flags().GetString("method")
			options.Method = method

			userAgent, _ := cmd.Flags().GetString("user-agent")
			options.UserAgent = userAgent

			cookies, _ := cmd.Flags().GetStringSlice("cookies")
			options.Cookies = cookies

			timeout, _ := cmd.Flags().GetDuration("timeout")
			options.Timeout = timeout

			form, _ := cmd.Flags().GetStringToString("form-item")
			formSlice, _ := cmd.Flags().GetStringSlice("form")
			for _, v := range formSlice {
				form[v] = ""
			}
			
			options.Form = form

			formType, _ := cmd.Flags().GetString("form-type")
			options.FormType = formType
			return ssti.Detect(url, options)
		},
	}
)

func init() {
	RootCmd.Flags().StringP("url", "u", "", "The target IP or domain to scan")
	RootCmd.MarkFlagRequired("url")

	RootCmd.Flags().StringP("method", "X", "GET", "The HTTP method to use")
	RootCmd.Flags().StringP("user-agent", "", "gossti 1.0.0", "Custom user-agent to use")
	RootCmd.Flags().StringSliceP("cookies", "C", nil, "Cookies to use (e.g. -C 'cookie1=value1; cookie2=value2')")
	RootCmd.Flags().Duration("timeout", 0, "Timeout for HTTP requests (e.g. 10s)")

	// Add support for multiple form fields
	RootCmd.Flags().StringToString("form-item", nil, "Form field to use (e.g. --form 'field1=value1' --form 'field2=value2')")
	RootCmd.Flags().StringSlice("form", nil, "Form fields to use (e.g. --form 'field1=value1,field2=value2')")
	RootCmd.Flags().String("form-type", "urlencoded", "Form type to use (e.g. urlencoded, multipart)")
}
