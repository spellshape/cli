package plugin_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	envtest "github.com/spellshape/cli/integration"
	pluginsconfig "github.com/spellshape/cli/spellshape/config/plugins"
	"github.com/spellshape/cli/spellshape/pkg/cmdrunner/step"
	"github.com/spellshape/cli/spellshape/services/plugin"
)

func TestAddRemovePlugin(t *testing.T) {
	var (
		require    = require.New(t)
		assert     = assert.New(t)
		env        = envtest.New(t)
		app        = env.Scaffold("github.com/test/blog")
		pluginRepo = "github.com/spellshape/example-plugin"

		assertPlugins = func(expectedLocalPlugins, expectedGlobalPlugins []pluginsconfig.Plugin) {
			localCfg, err := pluginsconfig.ParseDir(app.SourcePath())
			require.NoError(err)
			assert.ElementsMatch(expectedLocalPlugins, localCfg.Plugins, "unexpected local plugins")

			globalCfgPath, err := plugin.PluginsPath()
			require.NoError(err)
			globalCfg, err := pluginsconfig.ParseDir(globalCfgPath)
			require.NoError(err)
			assert.ElementsMatch(expectedGlobalPlugins, globalCfg.Plugins, "unexpected global plugins")
		}
	)

	// no plugins expected
	assertPlugins(nil, nil)

	env.Must(env.Exec("add plugin locally",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "plugin", "add", pluginRepo, "k1=v1", "k2=v2"),
			step.Workdir(app.SourcePath()),
		)),
	))

	// one local plugin expected
	assertPlugins(
		[]pluginsconfig.Plugin{
			{
				Path: pluginRepo,
				With: map[string]string{
					"k1": "v1",
					"k2": "v2",
				},
			},
		},
		nil,
	)

	env.Must(env.Exec("remove plugin locally",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "plugin", "remove", pluginRepo),
			step.Workdir(app.SourcePath()),
		)),
	))

	// no plugins expected
	assertPlugins(nil, nil)

	env.Must(env.Exec("add plugin globally",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "plugin", "add", pluginRepo, "-g"),
			step.Workdir(app.SourcePath()),
		)),
	))

	// one global plugins expected
	assertPlugins(
		nil,
		[]pluginsconfig.Plugin{
			{
				Path: pluginRepo,
			},
		},
	)

	env.Must(env.Exec("remove plugin globally",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "plugin", "remove", pluginRepo, "-g"),
			step.Workdir(app.SourcePath()),
		)),
	))

	// no plugins expected
	assertPlugins(nil, nil)
}

// TODO install network plugin test

func TestPluginScaffold(t *testing.T) {
	env := envtest.New(t)

	env.Must(env.Exec("add a plugin",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "plugin", "scaffold", "test"),
			step.Workdir(env.TmpDir()),
		)),
	))
}
