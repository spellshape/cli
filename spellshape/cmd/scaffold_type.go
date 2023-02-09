package spellshapecmd

import (
	"github.com/spf13/cobra"

	"github.com/spellshape/cli/spellshape/services/scaffolder"
)

// NewScaffoldType returns a new command to scaffold a type.
func NewScaffoldType() *cobra.Command {
	c := &cobra.Command{
		Use:     "type NAME [field]...",
		Short:   "Type definition",
		Args:    cobra.MinimumNArgs(1),
		PreRunE: gitChangesConfirmPreRunHandler,
		RunE:    scaffoldTypeHandler,
	}

	flagSetPath(c)
	flagSetClearCache(c)

	c.Flags().AddFlagSet(flagSetYes())
	c.Flags().AddFlagSet(flagSetScaffoldType())

	return c
}

func scaffoldTypeHandler(cmd *cobra.Command, args []string) error {
	return scaffoldType(cmd, args, scaffolder.DryType())
}
