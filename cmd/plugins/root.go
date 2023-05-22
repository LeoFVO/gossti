package plugins

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:     "plugins",
		Aliases: []string{},
		Short:   "Manage gossti plugins",
		Long:    `Update and manage gossti plugins`,
	}
)

func init() {
	RootCmd.AddCommand(updateCmd)
}