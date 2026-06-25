package processor

import (
	"image"
	"image/color"
)

func GrayScale(img image.Image) *image.RGBA {
	bound := img.Bounds()
	dst := image.NewRGBA(bound)

	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((299*r + 587*g + 114*b) / 1000 / 256)
			dst.Set(x, y, color.RGBA{
				R: gray,
				G: gray,
				B: gray,
				A: 255,
			})
		}
	}
	return dst
}
