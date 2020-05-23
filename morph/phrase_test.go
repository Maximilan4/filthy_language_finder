package morph

import (
	"reflect"
	"testing"
)

var sentence = "test sentence тест"

func TestPhrase_GetClearedRuString(t *testing.T) {
	phrase := StringToPhrase(sentence)
	if phrase.GetClearedRuString() != "тестсентенцтест" {
		t.Error()
	}
}

func TestPhrase_GetClearedEnString(t *testing.T) {
	phrase := StringToPhrase(sentence)
	if phrase.GetClearedEnString() != "testsentenc" {
		t.Error()
	}
}

func TestStringToPhrase(t *testing.T) {
	phrase := StringToPhrase(sentence)
	if phrase.originalString != sentence {
		t.Error()
	}
	if !reflect.DeepEqual(phrase.EnWords, []string{"test", "sentenc"}) {
		t.Error()
	}

	if !reflect.DeepEqual(phrase.RuWords, []string{"тест", "сентенц", "тест"}) {
		t.Error()
	}
}
