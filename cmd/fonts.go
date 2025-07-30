//go:build ignore

package main

import (
	"github.com/jphsd/fonts/alexbrush"
	"github.com/jphsd/fonts/cianan"
	"github.com/jphsd/fonts/xkcd"
	g2d "github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/color"
	"github.com/jphsd/graphics2d/image"
	"golang.org/x/image/font/sfnt"
)

type fi struct {
	name string
	data []byte
}

var (
	Fonts = []fi{
		{"alexbrush", alexbrush.OTF},
		{"cianan", cianan.TTF},
		{"xkcd", xkcd.TTF},
	}
)

// Draw a string using the fonts in the fonts package
func main() {
	str := "\"The time has come,\" the Walrus said, \"to speak of many things\""

	for _, f := range Fonts {
		width, height := 1000, 250
		img := image.NewRGBA(width, height, color.White)

		ttf, err := sfnt.Parse(f.data)
		if err != nil {
			panic(err)
		}

		// Create shape from string
		s, _, err := g2d.StringToShape(ttf, str)

		// Figure out where to place string
		bounds := s.Bounds()
		dx := bounds.Max.X - bounds.Min.X
		scale := float64(width) / float64(dx)
		xfm := g2d.Scale(scale, scale)
		s = s.Transform(xfm)
		bounds = s.Bounds()
		xfm = g2d.Translate(float64(-bounds.Min.X), float64(height/2))
		s = s.Transform(xfm)

		g2d.FillShape(img, s, g2d.BlackPen)

		// Capture image output
		image.SaveImage(img, f.name)
	}
}
