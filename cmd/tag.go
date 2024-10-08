package cmd

import (
	"errors"
	"g/pkg"

	"github.com/spf13/cobra"
)

var (
	create bool
	delete bool
	latest bool
)

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Short:   "Create/Remove tag git",
	Args: func(cmd *cobra.Command, args []string) error {
		if latest {
			return nil
		}
		if len(args) < 1 {
			return errors.New("not enough arguments, at least one should be specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if latest {
			latestTag()
			return
		}

		version := args[0]

		if create {
			createTag(version)
			return
		} else if delete {
			deleteTag(version)
			return
		}
		substituteTag(version)
	},
}

func substituteTag(version string) {
	deleteTag(version)
	createTag(version)
}

func createTag(version string) {
	pkg.RunCommand("git", "tag", version)
	pkg.RunCommand("git", "push", "origin", version)
}

func deleteTag(version string) {
	pkg.RunCommand("git", "tag", "-d", version)
	pkg.RunCommand("git", "push", "--delete", "origin", version)
}

func latestTag() {
	pkg.RunCommand("git", "describe", "--tags", "--abbrev=0")
}

func init() {
	tagCmd.Flags().BoolVarP(&create, "create", "c", false, "Create tag in local and repos")
	tagCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete tag from local and repos")
	tagCmd.Flags().BoolVarP(&latest, "latest", "l", false, "Show the latest tag")

	RootCmd.AddCommand(tagCmd)
}
