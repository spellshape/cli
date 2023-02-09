package app

import (
	"embed"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/plush/v4"

	"github.com/spellshape/cli/spellshape/pkg/cosmosgen"
	"github.com/spellshape/cli/spellshape/pkg/xgenny"
	"github.com/spellshape/cli/spellshape/templates/field/plushhelpers"
	"github.com/spellshape/cli/spellshape/templates/testutil"
)

//go:embed files/* files/**/*
var fs embed.FS

// New returns the generator to scaffold a new Cosmos SDK app.
func New(opts *Options) (*genny.Generator, error) {
	var (
		g        = genny.New()
		template = xgenny.NewEmbedWalker(fs, "files/", opts.AppPath)
	)
	if err := g.Box(template); err != nil {
		return g, err
	}
	ctx := plush.NewContext()
	ctx.Set("ModulePath", opts.ModulePath)
	ctx.Set("AppName", opts.AppName)
	ctx.Set("GitHubPath", opts.GitHubPath)
	ctx.Set("BinaryNamePrefix", opts.BinaryNamePrefix)
	ctx.Set("AddressPrefix", opts.AddressPrefix)
	ctx.Set("DepTools", cosmosgen.DepTools())

	plushhelpers.ExtendPlushContext(ctx)
	g.Transformer(xgenny.Transformer(ctx))
	g.Transformer(genny.Replace("{{appName}}", opts.AppName))
	g.Transformer(genny.Replace("{{binaryNamePrefix}}", opts.BinaryNamePrefix))

	// Create the 'testutil' package with the test helpers
	if err := testutil.Register(g, opts.AppPath); err != nil {
		return g, err
	}

	return g, nil
}
