package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/imageutil"
	"github.com/kylin419/CVforge/internal/morphology"
	"github.com/kylin419/CVforge/internal/progress"
)

type Dilation struct {
	Element morphology.Element
}

func (d Dilation) Process(img image.Image, reporter progress.Reporter) image.Image {
	gray := imageutil.Grayscale(img)
	return morphology.Dilate(gray, d.Element, reporter)
}
