package typed

import (
	"fmt"
	"strings"

	"github.com/spellshape/cli/spellshape/pkg/placeholder"
)

// ProtoGenesisStateMessage is the name of the proto message that represents the genesis state.
const ProtoGenesisStateMessage = "GenesisState"

// PatchGenesisTypeImport patches types/genesis.go content from the issue:
// https://github.com/spellshape/cli/issues/992
func PatchGenesisTypeImport(replacer placeholder.Replacer, content string) string {
	patternToCheck := "import ("
	replacement := fmt.Sprintf(`import (
%[1]v
)`, PlaceholderGenesisTypesImport)

	if !strings.Contains(content, patternToCheck) {
		content = replacer.Replace(content, PlaceholderGenesisTypesImport, replacement)
	}

	return content
}
