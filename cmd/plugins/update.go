package plugins

import (
	"github.com/LeoFVO/gossti/pkg/plugins"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Download and install the latest version of gossti plugins",
	Long:  `Download and install the latest version of gossti plugins`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := plugins.UpdatePlugins()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
}