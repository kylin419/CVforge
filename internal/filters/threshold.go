package filters

import (
	"image"
	"image/color"

	"github.com/kylin419/CVforge/internal/imageutil"
	"github.com/kylin419/CVforge/internal/progress"
)

type ThresholdType int

const (
	Binary ThresholdType = iota
	BinaryInv
	Truncate
	ToZero
	ToZeroInv
)

type Threshold struct {
	Value uint8
	Type  ThresholdType
}

func (t Threshold) Process(img image.Image, reporter progress.Reporter) image.Image {
	bound := img.Bounds()
	dst := image.NewGray(bound)
	gray := imageutil.Grayscale(img)
	reporter.Start(int64(bound.Dy()), "Threshold")
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			value := gray.GrayAt(x, y).Y
			var out uint8
			switch t.Type {
			case Binary:
				if value < uint8(t.Value) {
					out = 0
				} else {
					out = 255
				}
			case BinaryInv:
				if value >= t.Value {
					out = 0
				} else {
					out = 255
				}
			case Truncate:
				if value >= t.Value {
					out = t.Value
				} else {
					out = value
				}
			case ToZero:
				if value >= t.Value {
					out = value
				} else {
					out = 0
				}
			case ToZeroInv:
				if value >= t.Value {
					out = 0
				} else {
					out = value
				}
			}
			dst.SetGray(x, y, color.Gray{Y: out})
		}
		reporter.Update()
	}
	reporter.Update()
	reporter.Finish()
	return dst
}
