package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/squgeim/vocab-cli/dictionary"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	words, err := dictionary.ExtractAllWords()

	if err != nil {
		fmt.Println("Err!", err)
		os.Exit(1)
	}

	for _, w := range words {
		fmt.Println(w.Word)
	}
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the words in your collection",
	Args:  cobra.ExactArgs(0),
	Run:   list,
}
