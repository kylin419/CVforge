package pipeline

import (
	"image"

	"github.com/kylin419/CVforge/internal/filters"
)

type Pipeline struct {
	filters []filters.Filters
}

func New(fs ...filters.Filters) Pipeline {
	return Pipeline{
		filters: fs,
	}
}

func (p Pipeline) Run(img image.Image) image.Image {
	for _, f := range p.filters {
		img = f.Process(img)
	}
	return img
}
