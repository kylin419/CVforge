package filters

import (
	"image"
	"image/color"

	"github.com/kylin419/CVforge/internal/progress"
)

type Blur struct {
	Radius int
}

func (b Blur) Process(img image.Image, reporter progress.Reporter) image.Image {
	bound := img.Bounds()
	dst := image.NewRGBA(bound)
	if b.Radius < 1 {
		b.Radius = 1
	}
	if b.Radius > 10 {
		b.Radius = 10
	}
	reporter.Start(int64(bound.Dy()), "Blur")
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			var rSum uint32
			var gSum uint32
			var bSum uint32
			var count uint32
			for dy := -b.Radius; dy <= b.Radius; dy++ {
				for dx := -b.Radius; dx <= b.Radius; dx++ {
					nx := x + dx
					ny := y + dy

					if nx < bound.Min.X || nx >= bound.Max.X || ny < bound.Min.Y || ny >= bound.Max.Y {
						continue
					}
					r, g, b, _ := img.At(nx, ny).RGBA()
					rSum += r
					gSum += g
					bSum += b
					count++

				}
			}
			dst.Set(x, y, color.RGBA{
				R: uint8(rSum / count / 256),
				G: uint8(gSum / count / 256),
				B: uint8(bSum / count / 256),
				A: 255,
			})
		}
		reporter.Update()
	}
	reporter.Finish()
	return dst
}
