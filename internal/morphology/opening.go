package morphology

import (
	"image"

	"github.com/kylin419/CVforge/internal/progress"
)

func Opening(img *image.Gray, element Element, reporter progress.Reporter) *image.Gray {
	return Dilate(Erode(img, element, reporter), element, reporter)
}

func Closing(img *image.Gray, element Element, reporter progress.Reporter) *image.Gray {
	return Erode(Dilate(img, element, reporter), element, reporter)
}
