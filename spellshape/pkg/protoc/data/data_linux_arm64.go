package data

import _ "embed" // embed is required for binary embedding.

//go:embed protoc-linux-arm64
var binary []byte
