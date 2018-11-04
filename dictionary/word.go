package dictionary

import (
	"path/filepath"
)

// HasWord retruns if the given word exists in our dictionary
func HasWord(word string) bool {
	normalizedWord := utils.NormalizeWord(word)

	return utils.FileExists(pathToWord(normalizedWord))
}

func AddWord(word string) {

}

func pathToWord(word string) string {
	vocabDir := getVocabDir()

	return filepath.Join(vocabDir, word)
}
