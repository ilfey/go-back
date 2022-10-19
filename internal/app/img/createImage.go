package img

import (
	"fmt"

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
	W := float64(p.x)
	H := float64(p.y)

	ctx := gg.NewContext(p.x, p.y)

	// set background
	ctx.SetHexColor(p.bg)
	ctx.Clear()

	// set color and line width
	ctx.SetHexColor(p.fg)
	ctx.SetLineWidth(p.border)

	// create border
	ctx.DrawLine(0, 0, W, 0) // top
	ctx.DrawLine(W, 0, W, H) // right
	ctx.DrawLine(W, H, 0, H) // bottom
	ctx.DrawLine(0, H, 0, 0) // left

	// create x
	ctx.DrawLine(0, 0, W/4, H/4)     // top left
	ctx.DrawLine(W, 0, W-W/4, H/4)   // top right
	ctx.DrawLine(W, H, W-W/4, H-H/4) // bottom right
	ctx.DrawLine(0, H, W/4, H-H/4)   // bottom left

	// draw lines
	ctx.Stroke()

	// set text
	ctx.DrawStringAnchored(fmt.Sprintf("%dx%d", p.x, p.y), W/2, H/2, 0.5, 0.5)

	return ctx, nil
}
