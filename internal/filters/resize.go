package filters

import (
	"image"
	"time"

	"github.com/kylin419/CVforge/internal/progress"
)

type Resize struct {
	Width  int
	Height int
}

func (r Resize) Process(img image.Image, reporter progress.Reporter) image.Image {
	if r.Width <= 0 || r.Height <= 0 {
		return img
	}
	bound := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, r.Width, r.Height))
	reporter.Start(int64(r.Height), "Resize")
	for y := 0; y < r.Height; y++ {
		for x := 0; x < r.Width; x++ {
			imgX := x * bound.Dx() / r.Width
			imgY := y * bound.Dy() / r.Height
			c := img.At(imgX+bound.Min.X, imgY+bound.Min.Y)
			dst.Set(x, y, c)
		}
		reporter.Update()
		time.Sleep(time.Millisecond * 100)
	}
	reporter.Finish()
	return dst
}
