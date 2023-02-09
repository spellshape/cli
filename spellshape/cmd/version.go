package spellshapecmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/spellshape/cli/spellshape/version"
)

// NewVersion creates a new version command to show the Spellshape CLI version.
func NewVersion() *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Print the current build information",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println(version.Long(cmd.Context()))
		},
	}
	return c
}
