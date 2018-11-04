package dictionary

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

var currentUser *user.User

func init() {
	user, err := user.Current()

	if err != nil {
		fmt.Println("Cannot get hold of the current user")
		os.Exit(1)
	}

	currentUser = user
}

// IsInitialized returns if the dictionary has been initialized
func IsInitialized() bool {
	return dictionaryExists()
}

// Init initializes the vocab directory
func Init() error {
	if IsInitialized() {
		return nil
	}

	dir := getVocabDir()

	return os.MkdirAll(dir, os.ModePerm)
}

func getVocabDir() string {
	return filepath.Join(currentUser.HomeDir, ".vocab/dictionary")
}

func dictionaryExists() bool {
	vocabDir := getVocabDir()

	if _, err := os.Stat(vocabDir); os.IsNotExist(err) {
		return false
	}

	return true
}
