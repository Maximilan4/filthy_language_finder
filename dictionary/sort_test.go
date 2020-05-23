package dictionary

import (
	"reflect"
	"testing"
)

func TestDescStringSlice_Sort(t *testing.T) {
	words := descStringSlice{"a", "bc", "cvb"}
	expectedSlice := descStringSlice{"cvb", "bc", "a"}
	words.Sort()

	if !reflect.DeepEqual(expectedSlice, words) {
		t.Errorf("Expected %s, got %s", expectedSlice, words)
	}
}
