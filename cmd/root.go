package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/squgeim/vocab-cli/dictionary"
)

var rootCmd = &cobra.Command{
	Use:   "vocab",
	Short: "Vocab helps you integrate words into your vocabulary",
	Run:   main,
}

func init() {
	if dictionary.IsInitialized() {
		return
	}

	err := dictionary.Init()
	if err != nil {
		fmt.Println("Could not initialise the dictionary.")
		os.Exit(1)
	}
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
