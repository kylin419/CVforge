package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var equalizeCmd = &cobra.Command{
	Use:   "equalize input output",
	Short: "Histogram Equalization",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		pipe := pipeline.New(
			filters.HistogramEqualization{},
		)

		app := service.NewImageProcessor(pipe)

		if err := app.Process(args[0], args[1]); err != nil {
			return err
		}

		fmt.Println("\nDone")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(equalizeCmd)
}
