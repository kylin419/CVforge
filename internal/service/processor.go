package service

import (
	"github.com/kylin419/CVforge/internal/imageio"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/progress"
)

type ImageProcessor struct {
	pipeline pipeline.Pipeline
}

func NewImageProcessor(p pipeline.Pipeline) ImageProcessor {
	return ImageProcessor{
		pipeline: p,
	}
}

func (p ImageProcessor) Process(input, output string) error {
	img, err := imageio.Load(input)
	if err != nil {
		return err
	}
	bar := &progress.Bar{}
	result := p.pipeline.Run(img, bar)
	err = imageio.Save(output, result)
	return err
}
