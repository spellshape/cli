// Package scaffolder initializes Spellshape CLI apps and modifies existing ones
// to add more features in a later time.
package scaffolder

import (
	"context"
	"fmt"
	"path/filepath"

	chainconfig "github.com/spellshape/cli/spellshape/config/chain"
	"github.com/spellshape/cli/spellshape/pkg/cache"
	"github.com/spellshape/cli/spellshape/pkg/cosmosanalysis"
	"github.com/spellshape/cli/spellshape/pkg/cosmosgen"
	"github.com/spellshape/cli/spellshape/pkg/cosmosver"
	"github.com/spellshape/cli/spellshape/pkg/gocmd"
	"github.com/spellshape/cli/spellshape/pkg/gomodulepath"
)

// Scaffolder is Spellshape CLI app scaffolder.
type Scaffolder struct {
	// Version of the chain
	Version cosmosver.Version

	// path of the app.
	path string

	// modpath represents the go module path of the app.
	modpath gomodulepath.Path
}

// New creates a new scaffold app.
func New(appPath string) (Scaffolder, error) {
	sc, err := new(appPath)
	if err != nil {
		return sc, err
	}

	if sc.Version.LT(cosmosver.StargateFortyFourVersion) {
		return sc, fmt.Errorf(
			`⚠️ Your chain has been scaffolded with an old version of Cosmos SDK: %[1]v.
Please, follow the migration guide to upgrade your chain to the latest version:

https://docs.spellshape.com/migration`, sc.Version.String(),
		)
	}
	return sc, nil
}

// new creates a new scaffolder for an existent app.
func new(path string) (Scaffolder, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return Scaffolder{}, err
	}

	modpath, path, err := gomodulepath.Find(path)
	if err != nil {
		return Scaffolder{}, err
	}
	if err := cosmosanalysis.IsChainPath(path); err != nil {
		return Scaffolder{}, err
	}

	version, err := cosmosver.Detect(path)
	if err != nil {
		return Scaffolder{}, err
	}

	s := Scaffolder{
		Version: version,
		path:    path,
		modpath: modpath,
	}

	return s, nil
}

func finish(ctx context.Context, cacheStorage cache.Storage, path, gomodPath string) error {
	if err := protoc(ctx, cacheStorage, path, gomodPath); err != nil {
		return err
	}
	if err := gocmd.ModTidy(ctx, path); err != nil {
		return err
	}
	return gocmd.Fmt(ctx, path)
}

func protoc(ctx context.Context, cacheStorage cache.Storage, projectPath, gomodPath string) error {
	if err := cosmosgen.InstallDepTools(ctx, projectPath); err != nil {
		return err
	}

	confpath, err := chainconfig.LocateDefault(projectPath)
	if err != nil {
		return err
	}
	conf, err := chainconfig.ParseFile(confpath)
	if err != nil {
		return err
	}

	options := []cosmosgen.Option{
		cosmosgen.WithGoGeneration(gomodPath),
		cosmosgen.IncludeDirs(conf.Build.Proto.ThirdPartyPaths),
	}

	// Generate Typescript client code if it's enabled or when Vuex stores are generated
	if conf.Client.Typescript.Path != "" || conf.Client.Vuex.Path != "" { //nolint:staticcheck,nolintlint
		tsClientPath := chainconfig.TSClientPath(*conf)
		if !filepath.IsAbs(tsClientPath) {
			tsClientPath = filepath.Join(projectPath, tsClientPath)
		}

		options = append(options,
			cosmosgen.WithTSClientGeneration(
				cosmosgen.TypescriptModulePath(tsClientPath),
				tsClientPath,
				true,
			),
		)
	}

	if vuexPath := conf.Client.Vuex.Path; vuexPath != "" { //nolint:staticcheck,nolintlint
		if filepath.IsAbs(vuexPath) {
			vuexPath = filepath.Join(vuexPath, "generated")
		} else {
			vuexPath = filepath.Join(projectPath, vuexPath, "generated")
		}

		options = append(options,
			cosmosgen.WithVuexGeneration(
				cosmosgen.TypescriptModulePath(vuexPath),
				vuexPath,
			),
		)
	}

	if conf.Client.OpenAPI.Path != "" {
		openAPIPath := conf.Client.OpenAPI.Path
		if !filepath.IsAbs(openAPIPath) {
			openAPIPath = filepath.Join(projectPath, openAPIPath)
		}

		options = append(options, cosmosgen.WithOpenAPIGeneration(openAPIPath))
	}

	return cosmosgen.Generate(ctx, cacheStorage, projectPath, conf.Build.Proto.Path, options...)
}
