package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/imageutil"
	"github.com/kylin419/CVforge/internal/morphology"
	"github.com/kylin419/CVforge/internal/progress"
)

type Opening struct {
	Element morphology.Element
}

func (o Opening) Process(img image.Image, reporter progress.Reporter) image.Image {
	gray := imageutil.Grayscale(img)

	return morphology.Opening(gray, o.Element, reporter)
}
