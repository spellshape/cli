package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	spellshapecmd "github.com/spellshape/cli/spellshape/cmd"
	chainconfig "github.com/spellshape/cli/spellshape/config/chain"
	"github.com/spellshape/cli/spellshape/pkg/clictx"
	"github.com/spellshape/cli/spellshape/pkg/cliui/colors"
	"github.com/spellshape/cli/spellshape/pkg/cliui/icons"
	"github.com/spellshape/cli/spellshape/pkg/validation"
	"github.com/spellshape/cli/spellshape/pkg/xstrings"
)

func main() {
	os.Exit(run())
}

func run() int {
	const (
		exitCodeOK    = 0
		exitCodeError = 1
	)
	ctx := clictx.From(context.Background())

	cmd, cleanUp, err := spellshapecmd.New(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
		return exitCodeError
	}
	defer cleanUp()

	err = cmd.ExecuteContext(ctx)

	if errors.Is(ctx.Err(), context.Canceled) || errors.Is(err, context.Canceled) {
		fmt.Println("aborted")
		return exitCodeOK
	}

	if err != nil {
		var (
			validationErr validation.Error
			versionErr    chainconfig.VersionError
			msg           string
		)

		if errors.As(err, &validationErr) {
			msg = validationErr.ValidationInfo()
		} else {
			msg = err.Error()
		}

		// Make sure the error message starts with an upper case character
		msg = xstrings.ToUpperFirst(msg)

		fmt.Printf("%s %s\n", icons.NotOK, colors.Error(msg))

		if errors.As(err, &versionErr) {
			fmt.Println("Use a more recent CLI version or upgrade blockchain app's config")
		}

		return exitCodeError
	}
	return exitCodeOK
}
