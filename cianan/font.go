// cianan contains the Cianan font created by Feorag NicBhride from an 18th century Irish manuscript
package cianan

import (
	// force init for go:embed
	_ "embed"
)

//go:embed cianan.ttf
var TTF []byte
