package filters

import (
	"image"
	"image/color"
	"sort"

	"github.com/kylin419/CVforge/internal/progress"
)

type Median struct {
	Radius int
}

func median(value []int) uint8 {
	sort.Ints(value)
	mid := len(value) / 2
	return uint8(value[mid])
}

func (m Median) Process(img image.Image, reporter progress.Reporter) image.Image {
	bound := img.Bounds()
	dst := image.NewRGBA(bound)
	if m.Radius < 1 {
		m.Radius = 1
	}
	if m.Radius > 10 {
		m.Radius = 10
	}
	reporter.Start(int64(bound.Dy()), "Median Blur")
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			rValues := []int{}
			gValues := []int{}
			bValues := []int{}
			for dy := -m.Radius; dy <= m.Radius; dy++ {
				for dx := -m.Radius; dx <= m.Radius; dx++ {
					nx := x + dx
					ny := y + dy

					if nx < bound.Min.X ||
						nx >= bound.Max.X ||
						ny < bound.Min.Y ||
						ny >= bound.Max.Y {
						continue
					}
					r, g, b, _ := img.At(nx, ny).RGBA()

					rValues = append(rValues, int(r/256))
					gValues = append(gValues, int(g/256))
					bValues = append(bValues, int(b/256))
				}
				reporter.Update()
			}

			dst.Set(x, y, color.RGBA{
				R: uint8(median(rValues)),
				G: uint8(median(gValues)),
				B: uint8(median(bValues)),
				A: 255,
			})
			reporter.Update()
			reporter.Finish()
		}
	}
	return dst
}
