package morph

import (
	"github.com/gen1us2k/go-translit"
	"regexp"
	"strings"
)

type Phrase struct {
	originalString string
	EnWords[] string
	RuWords[] string
}

func (p *Phrase) GetClearedRuString() string {
	return strings.Join(p.RuWords, "")
}

func (p *Phrase) GetClearedEnString() string {
	return strings.Join(p.EnWords, "")
}

func StringToPhrase(sentence string) Phrase {
	phrase := Phrase{
		originalString: sentence,
		EnWords: make([]string, 0),
		RuWords: make([]string, 0),
	}
	phrase.loadSentence(sentence)
	return phrase
}

func (p *Phrase) loadSentence(sentence string) []string {
	reg, _ := regexp.Compile("[-()'\"#/@;^:&%±§<>{}_\x60+=~|.!?,[]]*")
	sentence = reg.ReplaceAllString(strings.ToLower(sentence), "")

	for _, word := range strings.Fields(sentence) {
		word = clearRepeatingRunes(word)
		if isEnWord(word) {
			word = clearEnWord(word)
			p.EnWords = append(p.EnWords, word)
			p.RuWords = append(p.RuWords, translit.Translit(word))
		} else if isRuWord(word) {
			p.RuWords = append(p.RuWords, clearRuWord(word))
		}

	}

	return strings.Fields(sentence)
}

