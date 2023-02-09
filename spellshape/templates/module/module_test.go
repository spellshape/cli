package module

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProtoPackageName(t *testing.T) {
	cases := []struct {
		name   string
		app    string
		module string
		want   string
	}{
		{
			name:   "name",
			app:    "spellshape",
			module: "test",
			want:   "spellshape.test",
		},
		{
			name:   "path",
			app:    "spellshape/cli",
			module: "test",
			want:   "spellshape.cli.test",
		},
		{
			name:   "path with dash",
			app:    "spellshape/c-li",
			module: "test",
			want:   "spellshape.cli.test",
		},
		{
			name:   "path with number prefix",
			app:    "0spellshape/cli",
			module: "test",
			want:   "_0spellshape.cli.test",
		},
		{
			name:   "path with number prefix and dash",
			app:    "0spellshape/cli",
			module: "test",
			want:   "_0spellshape.cli.test",
		},
		{
			name:   "module with dash",
			app:    "spellshape",
			module: "test-mod",
			want:   "spellshape.testmod",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, ProtoPackageName(tt.app, tt.module))
		})
	}
}
