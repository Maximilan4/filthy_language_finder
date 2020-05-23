package scoring

import (
	"github.com/Maximilan4/connor/dictionary"
	"github.com/Maximilan4/connor/morph"
	"strings"
)

type Scorer struct {
	ruDict *dictionary.Dictionary
	enDict *dictionary.Dictionary
}

type ScorerResult struct {
	Words []string
	Score int
}

func mergeScorerResults(sr1 ScorerResult, sr2 ScorerResult) ScorerResult{
	return ScorerResult{
		Words: append(sr1.Words, sr2.Words...),
		Score: sr1.Score + sr2.Score,
	}
}

func NewScorer(ruDict *dictionary.Dictionary, enDict *dictionary.Dictionary) *Scorer{
	return &Scorer{
		ruDict: ruDict,
		enDict: enDict,
	}
}

func (s *Scorer) ScoreAll(sentence string) ScorerResult {
	phrase := morph.StringToPhrase(sentence)
	scorerChan := make(chan ScorerResult)
	scoringFunc := func(scorerChan chan ScorerResult, sent string, dict *dictionary.Dictionary) {
		scorerChan <- score(sent, dict)
	}
	go scoringFunc(scorerChan, phrase.GetClearedRuString(), s.ruDict)
	go scoringFunc(scorerChan, phrase.GetClearedEnString(), s.enDict)

	return mergeScorerResults(<-scorerChan, <-scorerChan)
}

func score(sentence string, dict *dictionary.Dictionary) ScorerResult{
	found := search(sentence, dict)
	keys := make([]string, len(found))
	score := 0
	if len(found) == 0 {
		return ScorerResult{
			Words: keys,
			Score: score,
		}
	}

	i := 0
	for k, v := range found {
		keys[i] = k
		score += dict.GetWordScore(k) * v
		i++
	}

	return ScorerResult{
		Words: keys,
		Score: score,
	}
}

func search(sentence string, dict *dictionary.Dictionary) map[string]int {
	found := make(map[string]int)
	for _, dictWord := range dict.Words {
		sentLen := len(sentence)
		if len(dictWord) > sentLen {
			continue
		}

		if sentLen == 0 || sentLen < dict.MinLen {
			break
		}

		count := strings.Count(sentence, dictWord)
		if count == 0 {
			continue
		}

		found[dictWord] = count
		sentence = strings.ReplaceAll(sentence, dictWord, "")
	}

	return found
}
