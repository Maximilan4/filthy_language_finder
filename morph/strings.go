package morph

import (
	"bytes"
	"github.com/gen1us2k/go-translit"
	"github.com/kljensen/snowball/english"
	"github.com/kljensen/snowball/russian"
	"regexp"
	"strings"
	"unicode/utf8"
)

var ruRegex, _ = regexp.Compile("(?i)[А-ЯЁ]")
var enRegex, _ = regexp.Compile("(?i)[A-Z]")
var digitRegex, _ = regexp.Compile("[0-9]")
var numToEnMap = map[rune]rune{
	'0': 'o',
	'1': 'l',
	'2': 'z',
	'3': 'e',
	'4': 'f',
	'5': 's',
	'6': 'b',
	'7': 't',
	'8': 'x',
	'9': 'g',
}

var numToRuMap = map[rune]rune{
	'0': 'о',
	'1': -1,
	'2': -1,
	'3': 'е',
	'4': 'ч',
	'5': 'с',
	'6': 'б',
	'7': 'п',
	'8': 'х',
	'9': 'я',
}

func isRuWord(word string) bool {
	wordLen := utf8.RuneCountInString(word)
	if wordLen == 0 {
		return false
	}

	return len(ruRegex.FindAllString(word, -1)) >= wordLen/2
}

func isEnWord(word string) bool {
	wordLen := utf8.RuneCountInString(word)
	if wordLen == 0 {
		return false
	}

	return len(enRegex.FindAllString(word, -1)) >= wordLen/2
}

func hasRuRune(word string) bool {
	return ruRegex.MatchString(word)
}

func hasEnRune(word string) bool {
	return enRegex.MatchString(word)
}

func hasDigitRune(word string) bool {
	return digitRegex.MatchString(word)
}

func replaceNumByRuRune(word string) string {
	return strings.Map(func(r rune) rune {
		if replace, ok := numToRuMap[r]; ok {
			return replace
		}

		return r
	}, word)
}

func replaceNumByEnRune(word string) string {
	return strings.Map(func(r rune) rune {
		if replace, ok := numToEnMap[r]; ok {
			return replace
		}

		return r
	}, word)
}

func clearRuWord(word string) string {
	if hasDigitRune(word) {
		word = replaceNumByRuRune(word)
	}

	if hasEnRune(word) {
		word = translit.Translit(word)
	}

	return russian.Stem(word, true)
}

func clearEnWord(word string) string {
	if hasDigitRune(word) {
		word = replaceNumByEnRune(word)
	}

	return english.Stem(word, true)
}

func clearRepeatingRunes(word string) string {
	var buf bytes.Buffer
	var lastRune rune
	var repeatCount = 1
	for i, wordRune := range word {
		if wordRune != lastRune || i == 0 {
			repeatCount = 1
			lastRune = wordRune
		} else if wordRune == lastRune {
			repeatCount++
		}

		if repeatCount > 2 {
			continue
		}

		buf.WriteRune(wordRune)
	}

	return buf.String()
}
