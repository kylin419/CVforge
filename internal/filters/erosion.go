package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/imageutil"
	"github.com/kylin419/CVforge/internal/morphology"
	"github.com/kylin419/CVforge/internal/progress"
)

type Erosion struct {
	Element morphology.Element
}

func (e Erosion) Process(img image.Image, reporter progress.Reporter) image.Image {
	gray := imageutil.Grayscale(img)
	return morphology.Erode(gray, e.Element, reporter)
}
