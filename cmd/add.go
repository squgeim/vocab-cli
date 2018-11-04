package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/squgeim/vocab-cli/dictionary"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	word := args[0]

	if dictionary.HasWord(word) {
		fmt.Println("Already got the word. Forgot about it?")
		return
	}

	dictionary.AddWord(word)
}

var addCmd = &cobra.Command{
	Use:   "add [word]",
	Short: "Add a word to your collection",
	Args:  cobra.ExactArgs(1),
	Run:   add,
}
