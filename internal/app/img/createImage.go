package img

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type imageParams struct {
	x      int
	y      int
	bg     string
	fg     string
	tan    float64
	border int
}

func createImage(p *imageParams) (*image.RGBA, error) {
	bg, err := parseHexColor(p.bg)
	if err != nil {
		return nil, err
	}
	fg, err := parseHexColor(p.fg)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(0, 0, p.x, p.y))

	for yi := img.Bounds().Min.Y; yi < img.Bounds().Max.Y; yi++ {
		for xi := img.Bounds().Min.X; xi < img.Bounds().Max.X; xi++ {
			dx := int(float64(xi) * p.tan)
			cx := p.x / 2
			empty := (cx-cx/5 > xi || cx+cx/5 < xi)
			hb := p.border / 2

			//set background
			img.SetRGBA(xi, yi, bg)
			// create border
			if xi > p.x-p.border || xi < p.border || yi > p.y-p.border || yi < p.border {
				img.Set(xi, yi, fg)
			}
			// create x
			if sum := dx + yi; dx-hb < yi && dx+hb > yi && empty || sum > p.y-hb && sum < p.y+hb && empty {
				img.SetRGBA(xi, yi, fg)
			}
		}
	}

	// calc font size
	var fontSize float64
	if p.x/2 < p.y {
		fontSize = float64(p.x / 10)
	} else {
		fontSize = float64(p.y / 10)
	}

	// find font
	fontFile, err := findfont.Find("arial.ttf")
	if err != nil {
		for _, path := range findfont.List() {
			split := strings.Split(path, ".")
			if strings.ToLower(split[len(split)-1:][0]) == "ttf" {
				fontFile = path
				break
			}
		}
	}

	// read font file
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		logrus.Errorf("error while reading file: %s", fontFile)
	}

	fontLoaded, err := truetype.Parse([]byte(fontBytes))
	if err != nil {
		logrus.Error("font not supported")
		return img, nil
	}
	myFont := truetype.NewFace(fontLoaded, &truetype.Options{
		Size: fontSize,
	})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(fg),
		Face: myFont,
	}

	delta := d.MeasureString(fmt.Sprintf("%dx%d", p.x, p.y))
	d.Dot = fixed.Point26_6{
		X: fixed.I(p.x/2 - int(delta/fixed.I(2))),
		Y: fixed.I(p.y/2 + int(fontSize/float64(3))),
	}

	d.DrawString(fmt.Sprintf("%dx%d", p.x, p.y))

	return img, nil
}
