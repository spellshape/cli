package config

import (
	"github.com/spellshape/cli/spellshape/pkg/env"
	"github.com/spellshape/cli/spellshape/pkg/xfilepath"
)

// DirPath returns the path of configuration directory of Spellshape.
var DirPath = xfilepath.Mkdir(env.ConfigDir())
