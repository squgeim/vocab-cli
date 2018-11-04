package dictionary

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// ExtractAllWords returns a list of all the words in the dictionary
func ExtractAllWords() ([]Word, error) {
	var words []Word

	dir := getVocabDir()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		word, err := readWordFromFile(path)
		if err != nil {
			return err
		}

		words = append(words, word)

		return nil
	})

	return words, err
}

func readWordFromFile(path string) (Word, error) {
	var word Word
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return word, err
	}
	err = json.Unmarshal(b, &word)
	if err != nil {
		return word, err
	}

	return word, nil
}

// ExtractRandom returns a random word from the dictionary
func ExtractRandom() (Word, error) {
	files, err := ioutil.ReadDir(getVocabDir())

	if err != nil {
		return Word{}, err
	}

	wordPath := func() string {
		var info os.FileInfo

		for range files {
			i := files[rand.Intn(len(files))]

			if i.IsDir() {
				continue
			}

			info = i
			break
		}

		if info == nil {
			return ""
		}

		return filepath.Join(getVocabDir(), info.Name())
	}()

	if wordPath == "" {
		return Word{}, nil
	}

	word, err := readWordFromFile(wordPath)
	if err != nil {
		return Word{}, err
	}

	word.incrementAccess()

	return word, nil
}
