package service

import (
	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/imageio"
	"github.com/kylin419/CVforge/internal/pipeline"
)

type ImageProcessor struct {
	pipeline pipeline.Pipeline
}

func NewImageProcessor() ImageProcessor {
	pipe := pipeline.New(filters.Grayscale{})
	return ImageProcessor{
		pipeline: pipe,
	}
}

func (p ImageProcessor) Process(input, output string) error {
	img, err := imageio.Load(input)
	if err != nil {
		return err
	}
	result := p.pipeline.Run(img)
	err = imageio.Save(output, result)
	return err
}
