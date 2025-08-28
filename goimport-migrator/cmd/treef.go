package cmd

import (
	"github.com/Lunaryx-org/LunaPackage/goimport-migrator/shared"
	"github.com/spf13/cobra"
)

var treefCmd = &cobra.Command{
	Use:   "tree",
	Short: "Argument to print the files on the current workind directory",
	Run: func(cmd *cobra.Command, args []string) {
		shared.TreeFprint()
	},
}

func init() {
	rootCmd.AddCommand(treefCmd)
}
