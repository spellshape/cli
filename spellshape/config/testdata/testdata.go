package testdata

import (
	"testing"

	chainconfig "github.com/spellshape/cli/spellshape/config/chain"
	networkconfigTestdata "github.com/spellshape/cli/spellshape/config/chain/network/testdata"
	v0testdata "github.com/spellshape/cli/spellshape/config/chain/v0/testdata"
	v1testdata "github.com/spellshape/cli/spellshape/config/chain/v1/testdata"
	"github.com/spellshape/cli/spellshape/config/chain/version"
)

var Versions = map[version.Version][]byte{
	0: v0testdata.ConfigYAML,
	1: v1testdata.ConfigYAML,
}

var NetworkConfig = networkconfigTestdata.ConfigYAML

func GetLatestConfig(t *testing.T) *chainconfig.Config {
	return v1testdata.GetConfig(t)
}

func GetLatestNetworkConfig(t *testing.T) *chainconfig.Config {
	return networkconfigTestdata.GetConfig(t)
}
