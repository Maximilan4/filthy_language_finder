package messages

import (
	json2 "encoding/json"
)

type InputMessage struct {
	GameId int `json:"gameId"`
	Players []Player `json:"players"`
	ChatLog []LogMessage `json:"chatLog"`
}

type LogMessage struct {
	MessageId int `json:"messageId"`
	GameTime float64 `json:"gameTime"`
	ChatParticipantId int `json:"chatParticipantId"`
	Message string `json:"message"`
}

type Player struct {
	AccountId int `json:"accountId"`
	SummonerName string `json:"summonerName"`
	ChatParticipantId int `json:"chatParticipantId"`
}

func NewInputMessage(json []byte) (*InputMessage, error) {
	var inputMessage InputMessage
	err := json2.Unmarshal(json, &inputMessage)
	if err != nil {
		return &inputMessage, err
	}

	return &inputMessage, nil
}
