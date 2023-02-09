package query

import (
	"github.com/spellshape/cli/spellshape/pkg/multiformatname"
	"github.com/spellshape/cli/spellshape/templates/field"
)

// Options ...
type Options struct {
	AppName     string
	AppPath     string
	ModuleName  string
	ModulePath  string
	QueryName   multiformatname.Name
	Description string
	ResFields   field.Fields
	ReqFields   field.Fields
	Paginated   bool
}
