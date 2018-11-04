package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/squgeim/vocab-cli/dictionary"
)

func init() {
	rootCmd.AddCommand(randomCmd)
}

func random(cmd *cobra.Command, args []string) {
	word, err := dictionary.ExtractRandom()

	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	fmt.Println(word.Word)
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Extract a random word from your collection",
	Args:  cobra.ExactArgs(0),
	Run:   random,
}
