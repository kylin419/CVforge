package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var radius int

var medianCmd = &cobra.Command{
	Use:  "median input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pipe := pipeline.New(filters.Median{
			Radius: radius,
		})
		app := service.NewImageProcessor(pipe)
		if err := app.Process(args[0], args[1]); err != nil {
			return err
		}
		fmt.Println("\nDone")
		return nil
	},
}

func init() {
	medianCmd.Flags().IntVar(&radius, "radius", 3, "Median filter radius (window = 2*radius+1)")
	rootCmd.AddCommand(medianCmd)
}
