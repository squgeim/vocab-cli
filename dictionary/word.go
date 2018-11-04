package dictionary

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"time"
)

// HasWord retruns if the given word exists in our dictionary
func HasWord(word string) bool {
	normalizedWord := utils.NormalizeWord(word)

	return utils.FileExists(pathToWord(normalizedWord))
}

// AddWord adds the given word to the dictionary
func AddWord(word string) {
	w := Word{
		Word:          word,
		AddedOn:       time.Now(),
		AccessedTimes: []time.Time{time.Now()},
	}
	w.fetchAndAddMeaning()
	w.writeToFile()
}

func pathToWord(word string) string {
	vocabDir := getVocabDir()

	return filepath.Join(vocabDir, word)
}

// Word is the structure that defines a word in the dictionary
type Word struct {
	Word          string      `json:"word"`
	Meaning       string      `json:"meaning"`
	Usage         []Example   `json:"usage"`
	AddedOn       time.Time   `json:"addedOn"`
	AccessedTimes []time.Time `json:"accessedTimes"`
}

// Example is various example sentences
type Example struct {
	Sentence string
	source   string
}

func (w *Word) fetchAndAddMeaning() {

}

func (w *Word) writeToFile() {
	wordFile := pathToWord(utils.NormalizeWord(w.Word))

	b, err := json.Marshal(w)
	checkAndPanic(err)

	err = ioutil.WriteFile(wordFile, b, 0644)
	checkAndPanic(err)
}

func (w *Word) incrementAccess() {
	w.AccessedTimes = append(w.AccessedTimes, time.Now())
	w.writeToFile()
}
