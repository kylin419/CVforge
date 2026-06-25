package progress

import (
	"os"

	"github.com/schollz/progressbar/v3"
)

type Bar struct {
	bar *progressbar.ProgressBar
}

func (b *Bar) Start(
	total int64,
	desc string,
) {
	b.bar =
		progressbar.NewOptions64(total,
			progressbar.OptionSetDescription(desc),
			progressbar.OptionSetWriter(os.Stdout),
			progressbar.OptionShowCount(),
		)

}

func (b *Bar) Update() {
	if b.bar != nil {
		b.bar.Add(1)
	}

}

func (b *Bar) Finish() {

	if b.bar != nil {
		b.bar.Finish()
	}

}
