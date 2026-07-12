package filters

import (
	"image"

	"github.com/kylin419/CVforge/internal/kernel"
	"github.com/kylin419/CVforge/internal/progress"
)

type Gaussian struct {
	Size  int
	Sigma float64
}

func (g Gaussian) Process(img image.Image, reporter progress.Reporter) image.Image {
	reporter.Start(1, "Gaussian")
	if g.Size == 0 {
		g.Size = 3
	}
	if g.Sigma <= 0 {
		g.Sigma = 1
	}
	k := kernel.GaussianKernel(g.Size, g.Sigma)
	reporter.Update()
	reporter.Finish()
	return kernel.Apply(img, k)
}
