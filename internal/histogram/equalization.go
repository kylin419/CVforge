package histogram

import (
	"image"
	"image/color"
)

func Equalize(img *image.Gray) *image.Gray {
	bound := img.Bounds()
	dst := image.NewGray(bound)
	hist := Compute(img)
	cdf := ComputeCDF(hist)
	lut := BuildLUT(cdf, bound.Dx()*bound.Dy())

	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			old := img.GrayAt(x, y).Y
			new := lut[old]
			dst.SetGray(x, y, color.Gray{new})
		}
	}
	return dst
}
