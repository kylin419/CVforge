package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var resizeWidth, resizeHeight int

var resizeCmd = &cobra.Command{
	Use:  "resize input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pipe := pipeline.New(filters.Resize{Width: resizeWidth, Height: resizeHeight})
		app := service.NewImageProcessor(pipe)
		err := app.Process(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Println("\nDone")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resizeCmd)
	resizeCmd.Flags().IntVar(
		&resizeWidth,
		"w",
		800,
		"resize width",
	)
	resizeCmd.Flags().IntVar(
		&resizeHeight,
		"h",
		600,
		"resize height",
	)
}
