package dictionary

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Dictionary struct {
	Data   map[string] int
	Words  descStringSlice
	MinLen int
}

func (d *Dictionary) LoadFromFile(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		return false, err
	}
	if err := d.loadFromFile(filePath); err != nil {
		return false, err
	}

	return true, nil
}

func (d *Dictionary) loadFromFile(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		score, _ := strconv.ParseInt(line[1], 10, 64)
		word := strings.ToLower(line[0])
		d.Data[word] = int(score)
		d.Words = append(d.Words, word)
	}
	d.Words.Sort()
	d.findMinWordLen()

	return nil
}

func (d *Dictionary) findMinWordLen() {
	var buf int
	for _, word := range d.Words {
		wordLen := len(word)

		if buf == 0 || wordLen < buf {
			buf = wordLen
			continue
		}
	}

	d.MinLen = buf
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		Data:   map[string]int{},
		Words:  []string{},
		MinLen: 0,
	}
}
