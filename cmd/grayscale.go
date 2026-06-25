package cmd

import (
	"fmt"

	"github.com/kylin419/CVforge/internal/imageio"
	"github.com/kylin419/CVforge/internal/processor"
	"github.com/spf13/cobra"
)

var grayscaleCmd = &cobra.Command{
	Use:  "grayscale input output",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]
		output := args[1]
		img, err := imageio.Load(input)
		if err != nil {
			return err
		}
		result := processor.GrayScale(img)
		err = imageio.Save(output, result)
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
