package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GoSSTI",
	Short: "GoSSTI is a SSTI scanner for web application. Developed in Go.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verbose, err := cmd.Flags().GetCount("verbose")
		if err != nil {
			log.Fatal(err)
		}

		var level string 
		switch verbose {
			case 1:
				log.SetLevel(log.DebugLevel)
				level = "debug"
			case 2:
				log.SetLevel(log.TraceLevel)
				level = "trace"
			default:
						log.SetLevel(log.InfoLevel)
				level = "info"
		}
		log.Infof("Verbose level: %s\n", level)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringP("user-agent", "", "russti 1.0.0", "Custom user-agent to use")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Level of verbosity (can be repeated) from 0 to 2")	
}