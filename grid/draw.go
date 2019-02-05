package grid

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type PointsImage struct {
	Color func(interface{}) color.RGBA
}

func (p *PointsImage) Save(name string, g *DenseGrid) {
	img := image.NewRGBA(image.Rect(0, 0, g.Size.X, g.Size.Y))
	for iter, pt, ok := g.Size.FromOrigin().Start(); ok; pt, ok = iter.Next() {
		img.Set(pt.X, pt.Y, p.Color(g.Data[iter.Idx()]))
	}
	f, _ := os.OpenFile(name+".png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
