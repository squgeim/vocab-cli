package dictionary

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ExtractAllWords() ([]Word, error) {
	var words []Word

	dir := getVocabDir()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		var word Word

		if info.IsDir() {
			return nil
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &word)
		if err != nil {
			return err
		}
		words = append(words, word)

		return nil
	})

	return words, err
}
