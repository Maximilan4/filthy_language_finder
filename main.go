package main

import (
	"github.com/Maximilan4/connor/communication"
	"github.com/Maximilan4/connor/dictionary"
	"github.com/Maximilan4/connor/scoring"
	"github.com/Maximilan4/connor/settings"
	"github.com/joho/godotenv"
	"log"
)

type App struct {
	consumer *communication.Consumer
}

func NewApp() *App {
	return &App {
		consumer: communication.NewConsumer(
			settings.ConsumerQueue,
			settings.PublishQueue,
			settings.ErrorQueue),
	}
}


func (a *App) Run() {
	err := a.consumer.Connect(settings.AmqpConnectionSettings.GetConnectionUrl())
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Connected to broker")
	err = a.consumer.InitQueues()
	if err != nil {
		log.Fatal(err.Error())
	}

	handler, err := initHandler()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Starting consuming")
	err = a.consumer.Consume(handler)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func initHandler() (communication.MessageHandler, error) {
	dictRu := dictionary.NewDictionary()
	if _, err := dictRu.LoadFromFile(settings.RuDictionaryPath); err != nil {
		return nil, err
	}

	dictEn := dictionary.NewDictionary()
	if _, err := dictEn.LoadFromFile(settings.EnDictionaryPath); err != nil {
		return nil, err
	}
	scorer := scoring.NewMessageScorer(dictRu, dictEn)

	return communication.NewGameMessageHandler(scorer), nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := NewApp()
	app.Run()
}
