//go:build !relayer

package other_components_test

import (
	"testing"

	envtest "github.com/spellshape/cli/integration"
	"github.com/spellshape/cli/spellshape/pkg/cmdrunner/step"
)

func TestGenerateAnAppWithQuery(t *testing.T) {
	var (
		env = envtest.New(t)
		app = env.Scaffold("github.com/test/blog")
	)

	env.Must(env.Exec("create a query",
		step.NewSteps(step.New(
			step.Exec(
				envtest.SpellshapeApp,
				"s",
				"query",
				"--yes",
				"foo",
				"text",
				"vote:int",
				"like:bool",
				"-r",
				"foo,bar:int,foobar:bool",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a query with custom path",
		step.NewSteps(step.New(
			step.Exec(
				envtest.SpellshapeApp,
				"s",
				"query",
				"--yes",
				"AppPath",
				"text",
				"vote:int",
				"like:bool",
				"-r",
				"foo,bar:int,foobar:bool",
				"--path",
				"./blog",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a paginated query",
		step.NewSteps(step.New(
			step.Exec(
				envtest.SpellshapeApp,
				"s",
				"query",
				"--yes",
				"bar",
				"text",
				"vote:int",
				"like:bool",
				"-r",
				"foo,bar:int,foobar:bool",
				"--paginated",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a custom field type",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp,
				"s",
				"type",
				"--yes",
				"custom-type",
				"numInt:int",
				"numsInt:array.int",
				"numsIntAlias:ints",
				"numUint:uint",
				"numsUint:array.uint",
				"numsUintAlias:uints",
				"textString:string",
				"textStrings:array.string",
				"textStringsAlias:strings",
				"textCoin:coin",
				"textCoins:array.coin",
				"textCoinsAlias:coins",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a query with the custom field type as a response",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "s", "query", "--yes", "foobaz", "-r", "bar:CustomType"),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("should prevent using custom type in request params",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "s", "query", "--yes", "bur", "bar:CustomType"),
			step.Workdir(app.SourcePath()),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("create an empty query",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "s", "query", "--yes", "foobar"),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("should prevent creating an existing query",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "s", "query", "--yes", "foo", "bar"),
			step.Workdir(app.SourcePath()),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("create a module",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "s", "module", "--yes", "foo", "--require-registration"),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a query in a module",
		step.NewSteps(step.New(
			step.Exec(
				envtest.SpellshapeApp,
				"s",
				"query",
				"--yes",
				"foo",
				"text",
				"--module",
				"foo",
				"--desc",
				"foo bar foobar",
				"--response",
				"foo,bar:int,foobar:bool",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	app.EnsureSteady()
}
