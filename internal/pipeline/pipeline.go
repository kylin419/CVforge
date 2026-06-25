package pipeline

import (
	"image"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/progress"
)

type Pipeline struct {
	filters []filters.Filters
}

func New(fs ...filters.Filters) Pipeline {
	return Pipeline{
		filters: fs,
	}
}

func (p Pipeline) Run(img image.Image, reporter progress.Reporter) image.Image {
	for _, f := range p.filters {
		img = f.Process(img, reporter)
	}
	return img
}
