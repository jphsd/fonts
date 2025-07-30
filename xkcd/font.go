// xkcd contains the XKCD script font created from a sample of Randall Munroe's handwriting.
package xkcd

import (
	// force init for go:embed
	_ "embed"
)

//go:embed xkcd-script.ttf
var TTF []byte
