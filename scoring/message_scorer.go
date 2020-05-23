package scoring

import (
	"github.com/Maximilan4/connor/dictionary"
	"github.com/Maximilan4/connor/messages"
	"strings"
)

type MessageScorer struct {
	Scorer
}

type AccountScorerResult struct {
	ScorerResult
	AccountId int
}

type GameScorerResult struct {
	GameId int
	AccountsResult []AccountScorerResult
}

func NewMessageScorer(ruDict *dictionary.Dictionary, enDict *dictionary.Dictionary) *MessageScorer{
	return &MessageScorer{
		Scorer: *NewScorer(ruDict, enDict),
	}
}

func (ms *MessageScorer) ScoreGameMessages(gameMessage *messages.GameChatMessages) *GameScorerResult {
	scoringCount := len(gameMessage.Messages)
	results := GameScorerResult{GameId:gameMessage.GameId}

	resultsChan := make(chan AccountScorerResult, scoringCount)
	for _, playerMessages := range gameMessage.Messages {
		go ms.scoreMessage(playerMessages, resultsChan)
	}

	for i:=0; i < scoringCount; i++ {
		results.AccountsResult = append(results.AccountsResult, <-resultsChan)
	}

	return &results
}

func (ms *MessageScorer) scoreMessage(playerMessages messages.PlayerMessages, resultsChan chan AccountScorerResult) {
	msg := strings.Join(playerMessages.Messages, " ")

	result := AccountScorerResult{
		ScorerResult: ms.ScoreAll(msg),
		AccountId: playerMessages.AccountId,
	}

	resultsChan <- result
}
