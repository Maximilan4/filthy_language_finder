package communication

import (
	"encoding/json"
	"github.com/Maximilan4/connor/messages"
	"github.com/Maximilan4/connor/settings"
	"github.com/streadway/amqp"
	"log"
)

type Consuming interface {
	Connect(url settings.ConnectionUrl)
	InitQueues()
	Disconnect()
	Consume()
}

type Consumer struct {
	mainQueue    string
	publishQueue string
	errorQueue   string
	connection   *amqp.Connection
	channel      *amqp.Channel
}

func (c *Consumer) SendError(message *messages.ErrorOutMessage) {
	encoded, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.channel.Publish("", c.errorQueue, false, false, amqp.Publishing{
		Body: encoded,
	})

	log.Println("Error was sended to " + c.errorQueue)
}

func (c *Consumer) Consume(handler MessageHandler) error {
	deliveryChan, err := c.channel.Consume(c.mainQueue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for delivery := range deliveryChan {
		log.Println("Message received in " + c.mainQueue)
		c.channel.Ack(delivery.DeliveryTag, false)
		go func() {
			scorerResult, err := handler.HandleConsume(delivery.Body)
			if err != nil {
				c.SendError(handler.PrepareErrorMessage(err))
				return
			}

			publishMsg, err := json.Marshal(handler.PreparePublishMessage(scorerResult))
			if err != nil {
				c.SendError(handler.PrepareErrorMessage(err))
				return
			}

			err = c.channel.Publish("", c.publishQueue, false, false, amqp.Publishing{
				Body: publishMsg,
			})

			if err != nil {
				log.Fatal(err.Error())
			}

			log.Println("Message was sended to " + c.publishQueue)

		}()
	}

	return nil
}

func (c *Consumer) InitQueues() error {
	for _, queue := range []string{c.mainQueue, c.publishQueue, c.errorQueue} {
		_, err := c.channel.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Consumer) Disconnect() {
	c.channel.Close()
	c.connection.Close()
}

func (c *Consumer) Connect(url settings.ConnectionUrl) error {
	connection, err := amqp.Dial(string(url))
	if err != nil {
		return err
	}

	c.connection = connection
	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	c.channel = channel
	return nil
}

func NewConsumer(mainQueue string, publishQueue string, errorQueue string) *Consumer {
	return &Consumer{mainQueue: mainQueue, publishQueue: publishQueue, errorQueue: errorQueue}
}
