package morphology

import (
	"image"
	"image/color"

	"github.com/kylin419/CVforge/internal/progress"
)

func Dilate(img *image.Gray, element Element, reporter progress.Reporter) *image.Gray {
	if !element.Valid() {
		return img
	}
	bound := img.Bounds()
	dst := image.NewGray(bound)
	w := element.Width()
	h := element.Height()
	offsetX := w / 2
	offsetY := h / 2
	reporter.Start(int64(bound.Dy()), "Dilation")
	for y := bound.Min.Y; y < bound.Max.Y; y++ {

		for x := bound.Min.X; x < bound.Max.X; x++ {
			fill := false
		Check:
			for ey := 0; ey < h; ey++ {

				for ex := 0; ex < w; ex++ {
					nx := x + ex - offsetX
					ny := y + ey - offsetY
					if nx < bound.Min.X || nx >= bound.Max.X || ny < bound.Min.Y || ny >= bound.Max.Y {
						continue
					}
					if element.Data[ey][ex] == false {
						continue
					}
					pixel := img.GrayAt(nx, ny).Y
					if pixel == 255 {
						fill = true
						break Check
					}
				}

			}
			value := uint8(0)
			if fill {
				value = 255
			}
			dst.SetGray(x, y, color.Gray{Y: value})
		}
		reporter.Update()
	}
	reporter.Finish()
	return dst
}
