package filters

import (
	"image"
	"math"

	"github.com/kylin419/CVforge/internal/progress"
)

type Rotate struct {
	Angle float64
}

func (r Rotate) Process(img image.Image, reporter progress.Reporter) image.Image {
	bound := img.Bounds()
	w := bound.Dx()
	h := bound.Dy()

	rad := r.Angle * math.Pi / 180
	sin := math.Sin(rad)
	cos := math.Cos(rad)
	dst := image.NewRGBA(
		image.Rect(0, 0, w, h),
	)
	reporter.Start(int64(h), "Rotate")
	cx := w / 2
	cy := h / 2
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - cx
			dy := y - cy
			imgX := int(
				float64(dx)*cos+
					float64(dy)*sin,
			) + cx
			imgY := int(
				-float64(dx)*sin+
					float64(dy)*cos,
			) + cy
			if imgX >= 0 && imgX < w && imgY >= 0 && imgY < h {
				dst.Set(x, y, img.At(imgX+bound.Min.X, imgY+bound.Min.Y))
			}
		}
		reporter.Update()
	}
	reporter.Finish()
	return dst
}
