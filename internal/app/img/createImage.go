package img

import (
	"github.com/fogleman/gg"
)

type imageParams struct {
	x      int
	y      int
	bg     string
	fg     string
	tan    float64
	border float64
}

func createImage(p *imageParams) (*gg.Context, error) {

	ctx := gg.NewContext(p.x, p.y)

	// set background
	// create border
	// create x
	// set text

	return ctx, nil
}
