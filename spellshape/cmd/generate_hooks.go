package spellshapecmd

import (
	"github.com/spf13/cobra"

	"github.com/spellshape/cli/spellshape/pkg/cliui"
	"github.com/spellshape/cli/spellshape/pkg/cliui/icons"
	"github.com/spellshape/cli/spellshape/services/chain"
)

func NewGenerateHooks() *cobra.Command {
	c := &cobra.Command{
		Use:     "hooks",
		Short:   "TypeScript frontend client and React hooks",
		PreRunE: gitChangesConfirmPreRunHandler,
		RunE:    generateHooksHandler,
	}

	c.Flags().AddFlagSet(flagSetYes())
	c.Flags().StringP(flagOutput, "o", "", "React hooks output path")

	return c
}

func generateHooksHandler(cmd *cobra.Command, _ []string) error {
	session := cliui.New(cliui.StartSpinnerWithText(statusGenerating))
	defer session.End()

	c, err := newChainWithHomeFlags(
		cmd,
		chain.WithOutputer(session),
		chain.CollectEvents(session.EventBus()),
		chain.PrintGeneratedPaths())
	if err != nil {
		return err
	}

	cacheStorage, err := newCache(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString(flagOutput)
	if err != nil {
		return err
	}

	if err := c.Generate(cmd.Context(), cacheStorage, chain.GenerateHooks(output)); err != nil {
		return err
	}

	return session.Println(icons.OK, "Generated Typescript Client and React hooks")
}
