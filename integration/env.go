package envtest

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/stretchr/testify/require"

	"github.com/spellshape/cli/spellshape/pkg/cosmosfaucet"
	"github.com/spellshape/cli/spellshape/pkg/env"
	"github.com/spellshape/cli/spellshape/pkg/gocmd"
	"github.com/spellshape/cli/spellshape/pkg/gomodulepath"
	"github.com/spellshape/cli/spellshape/pkg/httpstatuschecker"
	"github.com/spellshape/cli/spellshape/pkg/xurl"
)

const (
	ConfigYML = "config.yml"
)

var (
	// SpellshapeApp hold the location of the spellshape binary used in the integration
	// tests. The binary is compiled the first time the env.New() function is
	// invoked.
	SpellshapeApp = path.Join(os.TempDir(), "spellshape-tests", "spellshape")

	isCI, _           = strconv.ParseBool(os.Getenv("CI"))
	compileBinaryOnce sync.Once
)

// Env provides an isolated testing environment and what's needed to
// make it possible.
type Env struct {
	t   *testing.T
	ctx context.Context
}

// New creates a new testing environment.
func New(t *testing.T) Env {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	e := Env{
		t:   t,
		ctx: ctx,
	}
	// To avoid conflicts with the default config folder located in $HOME, we
	// set an other one thanks to env var.
	cfgDir := path.Join(t.TempDir(), ".spellshape")
	env.SetConfigDir(cfgDir)

	t.Cleanup(cancel)
	compileBinaryOnce.Do(func() {
		compileBinary(ctx)
	})
	return e
}

func compileBinary(ctx context.Context) {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("unable to get working dir: %v", err))
	}
	_, appPath, err := gomodulepath.Find(wd)
	if err != nil {
		panic(fmt.Sprintf("unable to read go module path: %v", err))
	}
	var (
		output, binary = filepath.Split(SpellshapeApp)
		path           = path.Join(appPath, "spellshape", "cmd", "spellshape")
	)
	err = gocmd.BuildPath(ctx, output, binary, path, nil)
	if err != nil {
		panic(fmt.Sprintf("error while building binary: %v", err))
	}
}

func (e Env) T() *testing.T {
	return e.t
}

// SetCleanup registers a function to be called when the test (or subtest) and all its
// subtests complete.
func (e Env) SetCleanup(f func()) {
	e.t.Cleanup(f)
}

// Ctx returns parent context for the test suite to use for cancelations.
func (e Env) Ctx() context.Context {
	return e.ctx
}

// IsAppServed checks that app is served properly and servers are started to listening before ctx canceled.
func (e Env) IsAppServed(ctx context.Context, apiAddr string) error {
	checkAlive := func() error {
		addr, err := xurl.HTTP(apiAddr)
		if err != nil {
			return err
		}

		ok, err := httpstatuschecker.Check(ctx, fmt.Sprintf("%s/cosmos/base/tendermint/v1beta1/node_info", addr))
		if err == nil && !ok {
			err = errors.New("waiting for app")
		}
		if HasTestVerboseFlag() {
			fmt.Printf("IsAppServed at %s: %v\n", addr, err)
		}
		return err
	}

	return backoff.Retry(checkAlive, backoff.WithContext(backoff.NewConstantBackOff(time.Second), ctx))
}

// IsFaucetServed checks that faucet of the app is served properly.
func (e Env) IsFaucetServed(ctx context.Context, faucetClient cosmosfaucet.HTTPClient) error {
	checkAlive := func() error {
		_, err := faucetClient.FaucetInfo(ctx)
		return err
	}

	return backoff.Retry(checkAlive, backoff.WithContext(backoff.NewConstantBackOff(time.Second), ctx))
}

// TmpDir creates a new temporary directory.
func (e Env) TmpDir() (path string) {
	return e.t.TempDir()
}

// Home returns user's home dir.
func (e Env) Home() string {
	home, err := os.UserHomeDir()
	require.NoError(e.t, err)
	return home
}

// AppHome returns app's root home/data dir path.
func (e Env) AppHome(name string) string {
	return filepath.Join(e.Home(), fmt.Sprintf(".%s", name))
}

// Must fails the immediately if not ok.
// t.Fail() needs to be called for the failing tests before running Must().
func (e Env) Must(ok bool) {
	if !ok {
		e.t.FailNow()
	}
}

func (e Env) HasFailed() bool {
	return e.t.Failed()
}

func (e Env) RequireExpectations() {
	e.Must(e.HasFailed())
}

func Contains(s, partial string) bool {
	return strings.Contains(s, strings.TrimSpace(partial))
}

func HasTestVerboseFlag() bool {
	return flag.Lookup("test.v").Value.String() == "true"
}
