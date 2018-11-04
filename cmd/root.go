package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vocab",
	Short: "Vocab helps you integrate words into your vocabulary",
	Run:   main,
}

func main(cmd *cobra.Command, args []string) {
	fmt.Println("IN the main function")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
