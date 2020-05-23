// +build ignore

package main

import (
	"fmt"
	"github.com/Maximilan4/connor/settings"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

func main() {
	chatsDir := path.Join(settings.ResourcesPath, "chats")
	if _, err := os.Stat(chatsDir); err != nil {
		log.Fatal(err.Error())
	}

	files, err := ioutil.ReadDir(chatsDir)
	if err != nil {
		log.Fatal(err)
	}
	var group sync.WaitGroup

	group.Add(len(files))
	connection, _ := amqp.Dial(string(settings.AmqpConnectionSettings.GetConnectionUrl()))
	channel, _ := connection.Channel()

	for _, fileInfo := range files {
		if path.Ext(fileInfo.Name()) != ".json" {
			fmt.Println("skipped " + fileInfo.Name())
			group.Done()
			continue
		}

		filePath := path.Join(chatsDir, fileInfo.Name())
		go func() {
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err.Error())
			}
			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err.Error())
			}

			channel.Publish("", settings.ConsumerQueue, false, false, amqp.Publishing{
				Body: bytes,
			})

			defer file.Close()
			defer group.Done()
			fmt.Println("sended " + file.Name())
		}()
	}

	group.Wait()

}
