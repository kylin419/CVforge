package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/morphology"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var element string
var erosionCmd = &cobra.Command{
	Use:   "erosion input output",
	Short: "Apply morphological erosion",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		e, err := morphology.ParseElement(element)
		if err != nil {
			return err
		}
		pipe := pipeline.New(filters.Erosion{
			Element: e,
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
	rootCmd.AddCommand(erosionCmd)
	erosionCmd.Flags().StringVar(
		&element,
		"element",
		"square",
		"Structuring element (square|cross)",
	)
}
