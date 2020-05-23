package dictionary

import (
	"github.com/Maximilan4/connor/settings"
	"path"
	"testing"
)

func TestDictionary_LoadFromFile(t *testing.T) {
	dict := NewDictionary()
	if result, err := dict.LoadFromFile("not-exists"); result != false && err == nil {
		t.Errorf("Can`t catch error while loading not existing file")
	}
	dictPath := path.Join(settings.ResourcesPath, "tests", "test_dict.csv")

	if result, err := dict.LoadFromFile(dictPath); result == false && err != nil {
		t.Errorf("Error while loading test file %s", dictPath)
	}

	expected := 3
	if len(dict.Data) != expected  && len(dict.Words) != expected && dict.MinLen != expected {
		t.Errorf("Wrong dictionary load")
	}
}

func TestNewDictionary(t *testing.T) {
	dict := NewDictionary()
	if len(dict.Data) > 0 {
		t.Errorf("Default value of Dictionary.Data must be empty")
	}
	if len(dict.Words) > 0 {
		t.Errorf("Default value of Dictionary.Words must be empty")
	}
	if dict.MinLen != 0 {
		t.Errorf("Default value of Dictionary.MinLen must be 0")
	}
}
