package messages

type OutMessage struct {
	State string `json:"state"`
}

type PlayerResult struct {
	AccountId int `json:"accountId"`
	Score int `json:"score"`
}

type SuccessPayload struct {
	GameId int `json:"gameId"`
	Players [] PlayerResult `json:"players"`
}

type SuccessOutMessage struct {
	OutMessage
	Payload SuccessPayload `json:"payload"`
}

type ErrorPayload struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode int `json:"errorCode"`
}

type ErrorOutMessage struct {
	OutMessage
	Payload ErrorPayload `json:"payload"`
}
