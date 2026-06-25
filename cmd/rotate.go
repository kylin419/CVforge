package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var rotateAngle float64
var rotateCmd = &cobra.Command{
	Use:  "rotate input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pipe := pipeline.New(filters.Rotate{Angle: rotateAngle})
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
	rootCmd.AddCommand(rotateCmd)
	rotateCmd.Flags().Float64Var(&rotateAngle, "angle", 90, "Rotate angle")
}
