package dictionary

import (
	"fmt"
	"os"
	"strings"
)

type utilStruct struct{}

var utils utilStruct

func (u *utilStruct) NormalizeWord(word string) string {
	r := strings.NewReplacer(" ", "_")

	return r.Replace(strings.ToLower(word))
}

func (u *utilStruct) FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func checkAndPanic(err error) {
	if err != nil {
		fmt.Println("Ops")
		os.Exit(1)
	}
}
