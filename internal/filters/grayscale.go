package filters

import (
	"image"
	"image/color"

	"github.com/kylin419/CVforge/internal/progress"
)

type Grayscale struct{}

func (g Grayscale) Process(img image.Image, reporter progress.Reporter) image.Image {
	bound := img.Bounds()
	dst := image.NewRGBA(bound)
	reporter.Start(
		int64(bound.Dy()),
		"Grayscale",
	)
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
		reporter.Update()
	}
	reporter.Finish()
	return dst
}
