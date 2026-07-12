package filters

import (
	"github.com/kylin419/CVforge/internal/imageutil"
	"github.com/kylin419/CVforge/internal/kernel"
	"github.com/kylin419/CVforge/internal/progress"
	"image"
)

type Sobel struct{}

func (s Sobel) Process(img image.Image, reporter progress.Reporter) image.Image {
	reporter.Start(4, "Sobel")
	gray := imageutil.Grayscale(img)
	reporter.Update()
	gx := kernel.ApplyRaw(gray, kernel.SobelX)
	reporter.Update()
	gy := kernel.ApplyRaw(gray, kernel.SobelY)
	reporter.Update()
	gradient := kernel.CombineGradient(gx, gy)
	result := kernel.ToGrayImage(gradient)
	reporter.Update()

	reporter.Finish()
	return result
}
