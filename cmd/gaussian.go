package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var size int
var sigma float64

var gaussianCmd = &cobra.Command{
	Use:  "gaussian input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pipe := pipeline.New(filters.Gaussian{
			Size:  size,
			Sigma: sigma,
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
	gaussianCmd.Flags().IntVar(&size, "size", 3, "Gaussian kernel size (odd)")
	gaussianCmd.Flags().Float64Var(&sigma, "sigma", 1.0, "Gaussian sigma")
	rootCmd.AddCommand(gaussianCmd)
}
