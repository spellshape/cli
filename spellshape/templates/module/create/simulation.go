package modulecreate

import (
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/plush/v4"

	"github.com/spellshape/cli/spellshape/pkg/xgenny"
	"github.com/spellshape/cli/spellshape/templates/field"
	"github.com/spellshape/cli/spellshape/templates/field/plushhelpers"
)

// AddSimulation returns the generator to generate module_simulation.go file.
func AddSimulation(appPath, modulePath, moduleName string, params ...field.Field) (*genny.Generator, error) {
	var (
		g        = genny.New()
		template = xgenny.NewEmbedWalker(fsSimapp, "files/simapp/", appPath)
	)

	ctx := plush.NewContext()
	ctx.Set("moduleName", moduleName)
	ctx.Set("modulePath", modulePath)
	ctx.Set("params", params)

	plushhelpers.ExtendPlushContext(ctx)
	g.Transformer(genny.Replace("{{moduleName}}", moduleName))

	if err := xgenny.Box(g, template); err != nil {
		return nil, err
	}

	g.Transformer(xgenny.Transformer(ctx))
	return g, nil
}
