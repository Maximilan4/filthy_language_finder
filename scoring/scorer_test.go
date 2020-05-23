package scoring

import (
	"github.com/Maximilan4/connor/dictionary"
	"github.com/Maximilan4/connor/settings"
	"path"
	"reflect"
	"testing"
)

var dictPath = path.Join(settings.ResourcesPath, "tests", "test_dict.csv")
var dict1 = dictionary.NewDictionary()
var dict2 = dictionary.NewDictionary()
var scorer = NewScorer(dict1, dict2)

func TestScorer_ScoreAll(t *testing.T) {
	dict2.LoadFromFile(dictPath)
	expected := ScorerResult{
		Words: []string{"test", "word", "bad"},
		Score: 4,
	}

	fact := scorer.ScoreAll("test bad word very test")

	if expected.Score != fact.Score {
		t.Error()
	}

	founded := 0
	for _, expWord := range expected.Words {
		for _, factWord := range fact.Words {
			if expWord == factWord {
				founded++
				break
			}
		}
	}

	if founded != len(expected.Words) {
		t.Error()
	}
}

func TestMergeScoringResults(t *testing.T) {
	sr1 := ScorerResult{
		Words: []string{"test"},
		Score: 1,
	}
	sr2 := ScorerResult{
		Words: []string{"test2"},
		Score: 2,
	}
	expected := ScorerResult{
		Words: []string{"test", "test2"},
		Score: 3,
	}

	if !reflect.DeepEqual(expected, mergeScorerResults(sr1, sr2)) {
		t.Error()
	}
}

func TestNewScorer(t *testing.T) {
	dict1 := dictionary.NewDictionary()
	dict2 := dictionary.NewDictionary()
	expected := &Scorer{ruDict:dict1, enDict:dict2}
	expectedWrong := Scorer{ruDict:dict1, enDict:dict2}

	if !reflect.DeepEqual(expected, NewScorer(dict1, dict2)) {
		t.Error()
	}

	if reflect.DeepEqual(expectedWrong, NewScorer(dict1, dict2)) {
		t.Error()
	}
}
