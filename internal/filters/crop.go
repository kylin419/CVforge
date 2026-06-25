package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/progress"
)

type Crop struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (c Crop) Process(
	img image.Image,
	reporter progress.Reporter,
) image.Image {

	bound := img.Bounds()

	if c.Width <= 0 || c.Height <= 0 {
		return img
	}

	dst := image.NewRGBA(
		image.Rect(
			0,
			0,
			c.Width,
			c.Height,
		),
	)

	reporter.Start(
		int64(c.Height),
		"Crop",
	)

	for y := 0; y < c.Height; y++ {

		for x := 0; x < c.Width; x++ {

			imgX := c.X + x
			imgY := c.Y + y

			if imgX < bound.Min.X ||
				imgX >= bound.Max.X ||
				imgY < bound.Min.Y ||
				imgY >= bound.Max.Y {

				continue
			}

			dst.Set(
				x,
				y,
				img.At(
					imgX,
					imgY,
				),
			)
		}

		reporter.Update()
	}

	reporter.Finish()

	return dst
}
