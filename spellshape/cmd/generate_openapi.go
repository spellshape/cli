package spellshapecmd

import (
	"github.com/spf13/cobra"

	"github.com/spellshape/cli/spellshape/pkg/cliui"
	"github.com/spellshape/cli/spellshape/pkg/cliui/icons"
	"github.com/spellshape/cli/spellshape/services/chain"
)

func NewGenerateOpenAPI() *cobra.Command {
	c := &cobra.Command{
		Use:     "openapi",
		Short:   "OpenAPI spec for your chain",
		PreRunE: gitChangesConfirmPreRunHandler,
		RunE:    generateOpenAPIHandler,
	}

	c.Flags().AddFlagSet(flagSetYes())

	return c
}

func generateOpenAPIHandler(cmd *cobra.Command, _ []string) error {
	session := cliui.New(cliui.StartSpinnerWithText(statusGenerating))
	defer session.End()

	c, err := newChainWithHomeFlags(
		cmd,
		chain.WithOutputer(session),
		chain.CollectEvents(session.EventBus()),
		chain.PrintGeneratedPaths(),
	)
	if err != nil {
		return err
	}

	cacheStorage, err := newCache(cmd)
	if err != nil {
		return err
	}

	if err := c.Generate(cmd.Context(), cacheStorage, chain.GenerateOpenAPI()); err != nil {
		return err
	}

	return session.Println(icons.OK, "Generated OpenAPI spec")
}
