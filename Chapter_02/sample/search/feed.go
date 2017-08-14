package search

import (
	"os"
	"encoding/json"
	"fmt"
)

// Default data file name
const dataFile = "data.json"

// Default data dir name
const DefaultDir = "data"

// variable responsible to make the users be able
// to change the feed data file location
var dataDir string = DefaultDir

// Provide one way to the user to change the
// the feed data file location
func SetDataDirectory(dir string) {
	dataDir = dir
}

// The type that will be used into feed loading
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// Main function responsible to retrieve the feed data from the all registered matchers
func RetrieveFeeds() ([]*Feed, error) {

	// Opening the feed load file
	file, err := os.Open(fmt.Sprintf("%v/%v", dataDir, dataFile))
	if err != nil {
		return nil, err
	}

	// defering the closing of the data file
	defer file.Close()

	// Creating a feed array
	var feeds []*Feed

	// decoding the data file's content
	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		return nil, err
	}

	return feeds, err

}
