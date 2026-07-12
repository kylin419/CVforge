package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/histogram"
	"github.com/kylin419/CVforge/internal/imageutil"
	"github.com/kylin419/CVforge/internal/progress"
)

type HistogramEqualization struct{}

func (h HistogramEqualization) Process(img image.Image, reporter progress.Reporter) image.Image {
	reporter.Start(int64(img.Bounds().Dy()), "Equalize")
	gray := imageutil.Grayscale(img)
	reporter.Update()
	reporter.Finish()
	return histogram.Equalize(gray)
}
