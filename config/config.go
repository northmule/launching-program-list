package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type IntJson int
type PathJson string
type DaysJson []IntJson

type App struct {
	SortIndex IntJson  `json:"sortIndex"` // индекс сортировки
	Path      PathJson `json:"path"`      // Путь до программы, для запуска
	Days      DaysJson `json:"days"`      // Дни недели, по которым запускать (1-7)
}

type AppConfigs []App

const ConfigName = "apps.json"

func GetConfigs() (AppConfigs, error) {
	var configStruct AppConfigs
	fileData, _ := getFileData(ConfigName) // todo
	err := parseConfig(fileData, &configStruct)
	if err != nil {
		return nil, errors.New("config parsing error")
	}

	return configStruct, nil
}

func parseConfig(configJson []byte, config interface{}) error {
	var err = json.Unmarshal(configJson, &config)
	if err != nil {
		return err
	}

	return nil
}

func (s *IntJson) UnmarshalJSON(data []byte) error {
	data = bytes.Replace(data, []byte("\""), []byte(""), -1)
	value, _ := strconv.Atoi(string(data))
	*s = IntJson(value)
	return nil
}

func getFileData(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}
	file = bytes.Replace(file, []byte("\n"), []byte(""), -1)
	return file, nil
}

func testGetConfigData(fileName string) ([]byte, error) {
	jsonData := `[{"sortIndex":"3","path":"/usr/bin/kate","days":[1,2,3,4,5,6,7]},{"sortIndex":1,"path":"/snap/bin/telegram-desktop","days":[1,2,3,4,5,6,7]},{"sortIndex":2,"path":"/usr/bin/skypeforlinux","days":[1,2,3,4,5,6,7]}]`
	return bytes.Replace([]byte(jsonData), []byte("\n"), []byte(""), -1), nil
}
