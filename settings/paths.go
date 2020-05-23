package settings

import (
	"os"
	"path"
)

const ResourcesDirectoryName string = "connor_resources"

var RuDictionaryFileName = getEnvironment("RU_DICT_FILE_NAME", "dict_ru.csv")
var EnDictionaryFileName = getEnvironment("EN_DICT_FILE_NAME", "dict_en.csv")

var AppPath, _ = os.UserHomeDir()
var ResourcesPath = path.Join(AppPath, ResourcesDirectoryName)
var RuDictionaryPath = path.Join(ResourcesPath, RuDictionaryFileName)
var EnDictionaryPath = path.Join(ResourcesPath, EnDictionaryFileName)
