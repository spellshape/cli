//go:build !relayer

package chain_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	envtest "github.com/spellshape/cli/integration"
	"github.com/spellshape/cli/spellshape/pkg/cmdrunner/step"
	"github.com/spellshape/cli/spellshape/pkg/xos"
)

func TestServeWithWasm(t *testing.T) {
	t.Skip()

	var (
		env     = envtest.New(t)
		app     = env.Scaffold("github.com/test/sgblog")
		servers = app.RandomizeServerPorts()
	)

	env.Must(env.Exec("add Wasm module",
		step.NewSteps(step.New(
			step.Exec(envtest.SpellshapeApp, "s", "wasm", "--yes"),
			step.Workdir(app.SourcePath()),
		)),
	))

	var (
		ctx, cancel       = context.WithTimeout(env.Ctx(), envtest.ServeTimeout)
		isBackendAliveErr error
	)
	go func() {
		defer cancel()
		isBackendAliveErr = env.IsAppServed(ctx, servers.API)
	}()
	env.Must(app.Serve("should serve", envtest.ExecCtx(ctx)))

	require.NoError(t, isBackendAliveErr, "app cannot get online in time")
}

func TestServeWithCustomHome(t *testing.T) {
	var (
		env     = envtest.New(t)
		app     = env.Scaffold("github.com/test/sgblog2")
		servers = app.RandomizeServerPorts()
	)

	var (
		ctx, cancel       = context.WithTimeout(env.Ctx(), envtest.ServeTimeout)
		isBackendAliveErr error
	)
	go func() {
		defer cancel()
		isBackendAliveErr = env.IsAppServed(ctx, servers.API)
	}()
	env.Must(app.Serve("should serve", envtest.ExecCtx(ctx)))

	require.NoError(t, isBackendAliveErr, "app cannot get online in time")
}

func TestServeWithConfigHome(t *testing.T) {
	var (
		env     = envtest.New(t)
		app     = env.Scaffold("github.com/test/sgblog3")
		servers = app.RandomizeServerPorts()
	)

	var (
		ctx, cancel       = context.WithTimeout(env.Ctx(), envtest.ServeTimeout)
		isBackendAliveErr error
	)
	go func() {
		defer cancel()
		isBackendAliveErr = env.IsAppServed(ctx, servers.API)
	}()
	env.Must(app.Serve("should serve", envtest.ExecCtx(ctx)))

	require.NoError(t, isBackendAliveErr, "app cannot get online in time")
}

func TestServeWithCustomConfigFile(t *testing.T) {
	tmpDir := t.TempDir()

	var (
		env = envtest.New(t)
		app = env.Scaffold("github.com/test/sgblog4")
	)
	// Move config
	newConfig := "new_config.yml"
	newConfigPath := filepath.Join(tmpDir, newConfig)
	err := xos.Rename(filepath.Join(app.SourcePath(), "config.yml"), newConfigPath)
	require.NoError(t, err)
	app.SetConfigPath(newConfigPath)

	servers := app.RandomizeServerPorts()

	var (
		ctx, cancel       = context.WithTimeout(env.Ctx(), envtest.ServeTimeout)
		isBackendAliveErr error
	)
	go func() {
		defer cancel()
		isBackendAliveErr = env.IsAppServed(ctx, servers.API)
	}()
	env.Must(app.Serve("should serve", envtest.ExecCtx(ctx)))

	require.NoError(t, isBackendAliveErr, "app cannot get online in time")
}

// TestServeWithName tests serving a new chain scaffolded using a local name instead of a GitHub URI.
func TestServeWithName(t *testing.T) {
	var (
		env     = envtest.New(t)
		app     = env.Scaffold("sgblog5")
		servers = app.RandomizeServerPorts()
	)

	ctx, cancel := context.WithTimeout(env.Ctx(), envtest.ServeTimeout)

	var isBackendAliveErr error

	go func() {
		defer cancel()

		isBackendAliveErr = env.IsAppServed(ctx, servers.API)
	}()

	env.Must(app.Serve("should serve", envtest.ExecCtx(ctx)))

	require.NoError(t, isBackendAliveErr, "app cannot get online in time")
}
