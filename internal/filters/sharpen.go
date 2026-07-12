package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/kernel"
	"github.com/kylin419/CVforge/internal/progress"
)

type Sharpen struct{}

func (s Sharpen) Process(img image.Image, reporter progress.Reporter) image.Image {
	reporter.Start(1, "Sharpen")
	result := kernel.Apply(img, kernel.Sharpen)
	reporter.Update()
	reporter.Finish()
	return result
}
