package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var blurRadius int
var blurCmd = &cobra.Command{
	Use:  "blur input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pipe := pipeline.New(
			filters.Blur{
				Radius: blurRadius,
			},
		)
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
	rootCmd.AddCommand(blurCmd)

	blurCmd.Flags().
		IntVar(
			&blurRadius,
			"r",
			1,
			"blur radius",
		)
}
