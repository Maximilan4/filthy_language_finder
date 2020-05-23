package morph

import "testing"

func TestIsRuWord(t *testing.T) {
	if !isRuWord("тест") {
		t.Error()
	}

	if isRuWord("test") {
		t.Error()
	}

	if !isRuWord("teст") {
		t.Error()
	}
}

func TestIsEnWord(t *testing.T) {
	if isEnWord("тест") {
		t.Error()
	}

	if !isEnWord("test") {
		t.Error()
	}

	if !isEnWord("teст") {
		t.Error()
	}
}
