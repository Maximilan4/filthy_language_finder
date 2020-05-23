package communication

import (
	"github.com/Maximilan4/connor/messages"
	"github.com/Maximilan4/connor/scoring"
)

type MessageHandler interface {
	HandleConsume(msg []byte) (*scoring.GameScorerResult, error)
	PreparePublishMessage(result *scoring.GameScorerResult) *messages.SuccessOutMessage
	PrepareErrorMessage(error) *messages.ErrorOutMessage
}

type GameMessageHandler struct {
	Scorer *scoring.MessageScorer
}

func (gmh *GameMessageHandler) PreparePublishMessage(scoringResult *scoring.GameScorerResult) *messages.SuccessOutMessage {
	if len(scoringResult.AccountsResult) == 0 {
		return nil
	}

	out := &messages.SuccessOutMessage{}
	out.State = "success"
	out.Payload.GameId = scoringResult.GameId
	for _, accountResult := range scoringResult.AccountsResult {
		out.Payload.Players = append(out.Payload.Players, messages.PlayerResult{
			AccountId: accountResult.AccountId,
			Score:     accountResult.Score,
		})
	}

	return out
}

func (gmh *GameMessageHandler) PrepareErrorMessage(err error) *messages.ErrorOutMessage{
	out := &messages.ErrorOutMessage{}
	out.State = "error"
	out.Payload.ErrorMessage = err.Error()
	out.Payload.ErrorCode = 0

	return out
}

func (gmh *GameMessageHandler) HandleConsume(msg []byte) (*scoring.GameScorerResult, error) {
	input, err := messages.NewInputMessage(msg)
	if err != nil {
		return nil, err
	}

	gm := messages.NewGameChatMessages(input)

	return gmh.Scorer.ScoreGameMessages(gm), nil
}

func NewGameMessageHandler(scorer *scoring.MessageScorer) *GameMessageHandler {
	return &GameMessageHandler{Scorer: scorer}
}
