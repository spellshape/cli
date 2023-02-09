package spellshapecmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	pluginsconfig "github.com/spellshape/cli/spellshape/config/plugins"
)

func TestEnsureDefaultPlugins(t *testing.T) {
	tests := []struct {
		name                 string
		cfg                  *pluginsconfig.Config
		expectAddedInCommand bool
	}{
		{
			name:                 "should add because absent from config",
			cfg:                  &pluginsconfig.Config{},
			expectAddedInCommand: true,
		},
		{
			name: "should not add because already present in config",
			cfg: &pluginsconfig.Config{
				Plugins: []pluginsconfig.Plugin{{
					Path: "github.com/spellshape/cli-plugin-network@v42",
				}},
			},
			expectAddedInCommand: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{Use: "spellshape"}

			ensureDefaultPlugins(cmd, tt.cfg)

			expectedCmd := findCommandByPath(cmd, "spellshape network")
			if tt.expectAddedInCommand {
				assert.NotNil(t, expectedCmd)
			} else {
				assert.Nil(t, expectedCmd)
			}
		})
	}
}
