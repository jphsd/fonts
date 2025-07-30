// alexbrush contains the Alex Brush script font designed by Robert E Leuschke from the Open Font Library.
package alexbrush

import (
	// force init for go:embed
	_ "embed"
)

//go:embed alexbrush.otf
var OTF []byte
