package main

import (
	"connor/dictionary"
	"connor/settings"
	"fmt"
	"github.com/kljensen/snowball"
)

func main() {
	stemmed, err := snowball.Stem("бесполез", "russian", true)
	if err == nil{
		fmt.Println(stemmed)
	}

	dictRu := *dictionary.NewDictionary()
	if _, err := dictRu.LoadFromFile(settings.RuDictionaryPath); err != nil {
		fmt.Println(err.Error())
		return
	}

	dictEn := *dictionary.NewDictionary()
	if _, err := dictEn.LoadFromFile(settings.EnDictionaryPath); err != nil {
		fmt.Println(err.Error())
		return
	}
}

