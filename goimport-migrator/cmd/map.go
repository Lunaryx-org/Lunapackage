package cmd

import (
	"github.com/Lunaryx-org/LunaPackage/goimport-migrator/shared"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "map",
	Short: "Prints the number of files/folders under the working directory",
	Run: func(cmd *cobra.Command, args []string) {
		shared.Map()
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
