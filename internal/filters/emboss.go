package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/kernel"
	"github.com/kylin419/CVforge/internal/progress"
)

type Emboss struct{}

func (e Emboss) Process(img image.Image, reporter progress.Reporter) image.Image {
	reporter.Start(1, "Emboss")
	result := kernel.Apply(img, kernel.Emboss)
	reporter.Update()
	reporter.Finish()
	return result
}
