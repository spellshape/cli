package modulecreate

import (
	"fmt"
	"path/filepath"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/plush/v4"

	"github.com/spellshape/cli/spellshape/pkg/gomodulepath"
	"github.com/spellshape/cli/spellshape/pkg/placeholder"
	"github.com/spellshape/cli/spellshape/pkg/xgenny"
	"github.com/spellshape/cli/spellshape/templates/field/plushhelpers"
	"github.com/spellshape/cli/spellshape/templates/module"
)

const msgServiceImport = `"github.com/cosmos/cosmos-sdk/types/msgservice"`

// AddMsgServerConventionToLegacyModule add the files and the necessary modifications to an existing module that doesn't support MsgServer convention
// https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-031-msg-service.md
func AddMsgServerConventionToLegacyModule(replacer placeholder.Replacer, opts *MsgServerOptions) (*genny.Generator, error) {
	var (
		g        = genny.New()
		template = xgenny.NewEmbedWalker(fsMsgServer, "files/msgserver/", opts.AppPath)
	)

	g.RunFn(codecPath(replacer, opts.AppPath, opts.ModuleName))

	if err := g.Box(template); err != nil {
		return g, err
	}

	appModulePath := gomodulepath.ExtractAppPath(opts.ModulePath)

	ctx := plush.NewContext()
	ctx.Set("moduleName", opts.ModuleName)
	ctx.Set("modulePath", opts.ModulePath)
	ctx.Set("appName", opts.AppName)
	ctx.Set("protoPkgName", module.ProtoPackageName(appModulePath, opts.ModuleName))

	plushhelpers.ExtendPlushContext(ctx)
	g.Transformer(xgenny.Transformer(ctx))
	g.Transformer(genny.Replace("{{appName}}", opts.AppName))
	g.Transformer(genny.Replace("{{moduleName}}", opts.ModuleName))
	return g, nil
}

func codecPath(replacer placeholder.Replacer, appPath, moduleName string) genny.RunFn {
	return func(r *genny.Runner) error {
		path := filepath.Join(appPath, "x", moduleName, "types/codec.go")
		f, err := r.Disk.Find(path)
		if err != nil {
			return err
		}

		// Add msgservice import
		old := "import ("
		new := fmt.Sprintf(`%v
%v`, old, msgServiceImport)
		content := replacer.Replace(f.String(), old, new)

		// Add RegisterMsgServiceDesc method call
		template := `%[1]v

msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)`
		replacement := fmt.Sprintf(template, module.Placeholder3)
		content = replacer.Replace(content, module.Placeholder3, replacement)
		newFile := genny.NewFileS(path, content)
		return r.File(newFile)
	}
}
