package messages

type GameChatMessages struct {
	GameId int
	Messages []PlayerMessages
}

type PlayerMessages struct {
	AccountId int
	Messages []string
}

func (pm *PlayerMessages) addMessage(message string) {
	pm.Messages = append(pm.Messages, message)
}

func NewGameChatMessages(inputMessage *InputMessage) *GameChatMessages {
	playerMessagesMap := make(map[int]PlayerMessages)
	gcm := GameChatMessages{
		GameId: inputMessage.GameId,
	}

	for _, player := range inputMessage.Players {
		playerMessagesMap[player.ChatParticipantId] = PlayerMessages{
			AccountId: player.AccountId,
		}
	}

	for _, chatMsg := range inputMessage.ChatLog {
		pm := playerMessagesMap[chatMsg.ChatParticipantId]
		pm.addMessage(chatMsg.Message)
		playerMessagesMap[chatMsg.ChatParticipantId] = pm
	}

	for _, pm := range playerMessagesMap {
		gcm.Messages = append(gcm.Messages, pm)
	}

	return &gcm
}
