package cmd

import (
	"g/pkg"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push",
	Short:   "Push commit changes",
	Aliases: []string{"p"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		pushCommit(branch)
	},
}

func pushCommit(branch string) {
	pkg.RunCommand("git", "push", "origin", branch)
}

func init() {
	RootCmd.AddCommand(pushCmd)
}
