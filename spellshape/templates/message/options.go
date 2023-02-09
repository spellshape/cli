package message

import (
	"github.com/spellshape/cli/spellshape/pkg/multiformatname"
	"github.com/spellshape/cli/spellshape/templates/field"
)

// Options ...
type Options struct {
	AppName      string
	AppPath      string
	ModuleName   string
	ModulePath   string
	MsgName      multiformatname.Name
	MsgSigner    multiformatname.Name
	MsgDesc      string
	Fields       field.Fields
	ResFields    field.Fields
	NoSimulation bool
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	return nil
}
