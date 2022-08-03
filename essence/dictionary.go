package essence

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//Dictionary ....
type Dictionary struct {
	//Words represents the words in the dictionary
	Words map[string]struct{} `json:"words"`
}

//NewDictionary reads a file and builds a Dictionary object from it
func NewDictionary(file string) (*Dictionary, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	d := &Dictionary{
		Words: map[string]struct{}{},
	}

	for scanner.Scan() {
		inputWord := scanner.Text()
		_, ok := d.Words[inputWord]
		if !ok {
			d.Words[inputWord] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return d, nil
}

//GetPossibleMatches gets all possible matches from the dictionary object
func (d *Dictionary) GetPossibleMatches(prefix string) []string {
	matches := make([]string, 0)
	for w := range d.Words {
		if strings.HasPrefix(w, prefix) {
			matches = append(matches, w)
		}
	}

	return matches
}

//Update updates the dictionary with provided data
func (d *Dictionary) Update(words []string) {
	for _, w := range words {
		_, ok := d.Words[w]
		if !ok {
			d.Words[w] = struct{}{}
		} else {
			log.Printf("dictionary: word %v already exists in the dictionary.\n", w)
		}
	}
}

//WordCount return the words count in the dictionary
func (d *Dictionary) WordCount() int {
	return len(d.Words)
}
