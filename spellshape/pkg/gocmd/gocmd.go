package gocmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spellshape/cli/spellshape/pkg/cmdrunner/exec"
	"github.com/spellshape/cli/spellshape/pkg/cmdrunner/step"
	"github.com/spellshape/cli/spellshape/pkg/goenv"
)

const (
	// CommandInstall represents go "install" command.
	CommandInstall = "install"

	// CommandGet represents go "get" command.
	CommandGet = "get"

	// CommandBuild represents go "build" command.
	CommandBuild = "build"

	// CommandMod represents go "mod" command.
	CommandMod = "mod"

	// CommandModTidy represents go mod "tidy" command.
	CommandModTidy = "tidy"

	// CommandModVerify represents go mod "verify" command.
	CommandModVerify = "verify"

	// CommandFmt represents go "fmt" command.
	CommandFmt = "fmt"

	// CommandEnv represents go "env" command.
	CommandEnv = "env"

	// EnvGOARCH represents GOARCH variable.
	EnvGOARCH = "GOARCH"
	// EnvGOMOD represents GOMOD variable.
	EnvGOMOD = "GOMOD"
	// EnvGOOS represents GOOS variable.
	EnvGOOS = "GOOS"

	// FlagGcflags represents gcflags go flag.
	FlagGcflags = "-gcflags"
	// FlagGcflagsValueDebug represents debug go flags.
	FlagGcflagsValueDebug = "all=-N -l"
	// FlagLdflags represents ldflags go flag.
	FlagLdflags = "-ldflags"
	// FlagMod represents mod go flag.
	FlagMod = "-mod"
	// FlagModValueReadOnly represents readonly go flag.
	FlagModValueReadOnly = "readonly"
	// FlagOut represents out go flag.
	FlagOut = "-o"
)

// Env returns the value of `go env name`.
func Env(name string) (string, error) {
	var b bytes.Buffer
	err := exec.Exec(context.Background(), []string{Name(), CommandEnv, name}, exec.StepOption(step.Stdout(&b)))
	return b.String(), err
}

// Name returns the name of Go binary to use.
func Name() string {
	custom := os.Getenv("GONAME")
	if custom != "" {
		return custom
	}
	return "go"
}

// Fmt runs go fmt on path.
func Fmt(ctx context.Context, path string, options ...exec.Option) error {
	return exec.Exec(ctx, []string{Name(), CommandFmt, "./..."}, append(options, exec.StepOption(step.Workdir(path)))...)
}

// ModTidy runs go mod tidy on path with options.
func ModTidy(ctx context.Context, path string, options ...exec.Option) error {
	return exec.Exec(ctx, []string{Name(), CommandMod, CommandModTidy}, append(options, exec.StepOption(step.Workdir(path)))...)
}

// ModVerify runs go mod verify on path with options.
func ModVerify(ctx context.Context, path string, options ...exec.Option) error {
	return exec.Exec(ctx, []string{Name(), CommandMod, CommandModVerify}, append(options, exec.StepOption(step.Workdir(path)))...)
}

// BuildPath runs go install on cmd folder with options.
func BuildPath(ctx context.Context, output, binary, path string, flags []string, options ...exec.Option) error {
	binaryOutput, err := binaryPath(output, binary)
	if err != nil {
		return err
	}
	command := []string{
		Name(),
		CommandBuild,
		FlagOut, binaryOutput,
	}
	command = append(command, flags...)
	command = append(command, ".")
	return exec.Exec(ctx, command, append(options, exec.StepOption(step.Workdir(path)))...)
}

// Build runs go build on path with options.
func Build(ctx context.Context, out, path string, flags []string, options ...exec.Option) error {
	command := []string{
		Name(),
		CommandBuild,
		FlagOut, out,
	}
	command = append(command, flags...)
	return exec.Exec(ctx, command, append(options, exec.StepOption(step.Workdir(path)))...)
}

// InstallAll runs go install ./... on path with options.
func InstallAll(ctx context.Context, path string, flags []string, options ...exec.Option) error {
	command := []string{
		Name(),
		CommandInstall,
	}
	command = append(command, flags...)
	command = append(command, "./...")
	return exec.Exec(ctx, command, append(options, exec.StepOption(step.Workdir(path)))...)
}

// Install runs go install pkgs on path with options.
func Install(ctx context.Context, path string, pkgs []string, options ...exec.Option) error {
	command := []string{
		Name(),
		CommandInstall,
	}
	command = append(command, pkgs...)
	return exec.Exec(ctx, command, append(options, exec.StepOption(step.Workdir(path)))...)
}

// Get runs go get pkgs on path with options.
func Get(ctx context.Context, path string, pkgs []string, options ...exec.Option) error {
	command := []string{
		Name(),
		CommandGet,
	}
	command = append(command, pkgs...)
	return exec.Exec(ctx, command, append(options, exec.StepOption(step.Workdir(path)))...)
}

// Ldflags returns a combined ldflags set from flags.
func Ldflags(flags ...string) string {
	return strings.Join(flags, " ")
}

// BuildTarget builds a GOOS:GOARCH pair.
func BuildTarget(goos, goarch string) string {
	return fmt.Sprintf("%s:%s", goos, goarch)
}

// ParseTarget parses GOOS:GOARCH pair.
func ParseTarget(t string) (goos, goarch string, err error) {
	parsed := strings.Split(t, ":")
	if len(parsed) != 2 {
		return "", "", errors.New("invalid Go target, expected in GOOS:GOARCH format")
	}

	return parsed[0], parsed[1], nil
}

// PackageLiteral returns the string representation of package part of go get [package].
func PackageLiteral(path, version string) string {
	return fmt.Sprintf("%s@%s", path, version)
}

// binaryPath determines the path where binary will be located at.
func binaryPath(output, binary string) (string, error) {
	if output != "" {
		outputAbs, err := filepath.Abs(output)
		if err != nil {
			return "", err
		}
		return filepath.Join(outputAbs, binary), nil
	}
	return filepath.Join(goenv.Bin(), binary), nil
}
