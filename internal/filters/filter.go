package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/progress"
)

type Filters interface {
	Process(img image.Image, reporter progress.Reporter) image.Image
}
