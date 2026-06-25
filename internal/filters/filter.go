package filters

import "image"

type Filters interface {
	Process(img image.Image) image.Image
}
