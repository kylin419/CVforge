package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/imageutil"
	"github.com/kylin419/CVforge/internal/morphology"
	"github.com/kylin419/CVforge/internal/progress"
)

type Closing struct {
	Element morphology.Element
}

func (c Closing) Process(img image.Image, reporter progress.Reporter) image.Image {
	gray := imageutil.Grayscale(img)

	return morphology.Closing(
		gray,
		c.Element,
		reporter,
	)
}
