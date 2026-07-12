package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/filters"
	"github.com/kylin419/CVforge/internal/pipeline"
	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var value uint8
var thresholdType string

var thresholdCmd = &cobra.Command{
	Use:  "threshold input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := filters.ParseThresholdType(thresholdType)
		if err != nil {
			return err
		}
		pipe := pipeline.New(filters.Threshold{
			Value: value,
			Type:  t,
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
	thresholdCmd.Flags().StringVar(&thresholdType, "type", "binary", "Threshold type (binary, binary-inv, truncate, tozero, tozero-inv)")
	thresholdCmd.Flags().Uint8Var(&value, "value", 127, "Threshold value")
	rootCmd.AddCommand(thresholdCmd)
}
