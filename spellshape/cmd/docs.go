package spellshapecmd

import (
	"github.com/spf13/cobra"

	"github.com/spellshape/cli/docs"
	"github.com/spellshape/cli/spellshape/pkg/localfs"
	"github.com/spellshape/cli/spellshape/pkg/markdownviewer"
)

func NewDocs() *cobra.Command {
	c := &cobra.Command{
		Use:   "docs",
		Short: "Show Spellshape CLI docs",
		Args:  cobra.NoArgs,
		RunE:  docsHandler,
	}
	return c
}

func docsHandler(*cobra.Command, []string) error {
	path, cleanup, err := localfs.SaveTemp(docs.Docs)
	if err != nil {
		return err
	}
	defer cleanup()

	return markdownviewer.View(path)
}
