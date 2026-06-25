package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/service"
	"github.com/spf13/cobra"
)

var grayscaleCmd = &cobra.Command{
	Use:  "grayscale input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]
		output := args[1]
		app := service.NewImageProcessor()
		err := app.Process(input, output)
		if err != nil {
			return err
		}
		fmt.Println("Done")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(grayscaleCmd)
}
