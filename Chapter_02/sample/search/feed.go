package search

import (
	"os"
	"encoding/json"
	"fmt"
)

const dataFile = "data.json"
const DefaultDir = "data"

var dataDir string = DefaultDir

func SetDataDirectory(dir string) {
	dataDir = dir
}

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {

	file, err := os.Open(fmt.Sprintf("%v/%v", dataDir, dataFile))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed

	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		return nil, err
	}

	return feeds, err

}
