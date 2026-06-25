package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var (
	cropX int
	cropY int
	cropW int
	cropH int
)

var cropCmd = &cobra.Command{

	Use: "crop input output",

	Args: cobra.ExactArgs(2),

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		pipe := pipeline.New(
			filters.Crop{
				X:      cropX,
				Y:      cropY,
				Width:  cropW,
				Height: cropH,
			},
		)

		app := service.NewImageProcessor(pipe)

		err := app.Process(
			args[0],
			args[1],
		)

		if err != nil {
			return err
		}

		fmt.Println("\nDone")

		return nil
	},
}

func init() {

	rootCmd.AddCommand(cropCmd)

	cropCmd.Flags().
		IntVar(
			&cropX,
			"x",
			0,
			"crop x",
		)

	cropCmd.Flags().
		IntVar(
			&cropY,
			"y",
			0,
			"crop y",
		)

	cropCmd.Flags().
		IntVar(
			&cropW,
			"w",
			100,
			"crop width",
		)

	cropCmd.Flags().
		IntVar(
			&cropH,
			"h",
			100,
			"crop height",
		)
}
