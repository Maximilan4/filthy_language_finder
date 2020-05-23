// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"github.com/Maximilan4/connor/messages"
	"github.com/Maximilan4/connor/settings"
	"github.com/streadway/amqp"
	"strconv"
)

func main() {
	connection, _ := amqp.Dial(string(settings.AmqpConnectionSettings.GetConnectionUrl()))
	channel, _ := connection.Channel()
	deliveryChan, _ := channel.Consume(settings.PublishQueue, "", true, false, false, false, amqp.Table{})

	for delivery := range deliveryChan {
		msg := messages.SuccessOutMessage{}
		json.Unmarshal(delivery.Body, &msg)
		for _, player := range msg.Payload.Players {
			fmt.Println("User :" + strconv.Itoa(player.AccountId))
			fmt.Println("Score :" + strconv.Itoa(player.Score))
		}
	}
}
