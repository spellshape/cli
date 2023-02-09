// this tool generates Spellshape CLI docs to be placed in the docs/cli dir and deployed
// on docs.spellshape.com
package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	spellshapecmd "github.com/spellshape/cli/spellshape/cmd"
	pluginsconfig "github.com/spellshape/cli/spellshape/config/plugins"
	"github.com/spellshape/cli/spellshape/pkg/env"
	"github.com/spellshape/cli/spellshape/services/plugin"
)

const head = `---
description: Spellshape CLI docs.
---

# CLI commands

Documentation for Spellshape CLI.
`

func main() {
	outPath := flag.String("out", ".", ".md file path to place Spellshape CLI docs inside")
	flag.Parse()
	if err := run(*outPath); err != nil {
		log.Fatal(err)
	}
}

func run(outPath string) error {
	// We want to have documentation for commands that are implemented in plugins.
	// To do that, we need to add the related plugins in the config.
	// To avoid conflicts with user config, set an alternate config dir in tmp.
	dir, err := os.MkdirTemp("", ".spellshape")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	env.SetConfigDir(dir)
	pluginDir, err := plugin.PluginsPath()
	if err != nil {
		return err
	}
	cfg, err := pluginsconfig.ParseDir(pluginDir)
	if err != nil {
		return err
	}
	cfg.Plugins = append(cfg.Plugins, pluginsconfig.Plugin{
		// Add network plugin
		Path: spellshapecmd.PluginNetworkPath,
	})
	if err := cfg.Save(); err != nil {
		return err
	}

	cmd, cleanUp, err := spellshapecmd.New(context.Background())
	if err != nil {
		return err
	}
	defer cleanUp()

	// Run ExecuteC so cobra adds the completion command.
	cmd, _ = cmd.ExecuteC()

	return generate(cmd, outPath)
}

func generate(cmd *cobra.Command, outPath string) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := fmt.Fprint(f, head); err != nil {
		return err
	}

	return generateCmd(cmd, f)
}

func generateCmd(cmd *cobra.Command, w io.Writer) error {
	cmd.DisableAutoGenTag = true

	b := &bytes.Buffer{}
	if err := doc.GenMarkdownCustom(cmd, b, linkHandler); err != nil {
		return err
	}

	// here we change sub titles to bold styling. Otherwise, these titles will get
	// printed in the right menu of docs.spellshape.network which is unpleasant because
	// we only want to see a list of all available commands without the extra noise.
	sc := bufio.NewScanner(b)
	for sc.Scan() {
		t := sc.Text()
		if strings.HasPrefix(t, "###") {
			t = strings.TrimPrefix(t, "### ")
			t = fmt.Sprintf("**%s**", t)
		}
		if _, err := fmt.Fprintln(w, t); err != nil {
			return err
		}
	}

	for _, cmd := range cmd.Commands() {
		if !cmd.IsAvailableCommand() || cmd.IsAdditionalHelpTopicCommand() {
			continue
		}

		io.WriteString(w, "\n")

		if err := generateCmd(cmd, w); err != nil {
			return err
		}
	}

	return nil
}

func linkHandler(link string) string {
	link = strings.ReplaceAll(link, "_", "-")
	link = strings.TrimSuffix(link, ".md")
	link = "#" + link
	return link
}
