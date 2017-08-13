package search

import (
	"testing"

	"fmt"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestRetrieveFeedsInvalidDataDir(t *testing.T) {

	invalidDir := "rldfkvmdflkv"
	SetDataDirectory(invalidDir)

	_, err := RetrieveFeeds()

	if err == nil {
		t.Fatal(fmt.Sprintf("should cannot load the %v on %v", dataFile, invalidDir), ballotX)
	}

	t.Log("feeds haven't loaded correctly : ", err, checkMark)
}

func TestRetrieveFeedsValidDataDir(t *testing.T) {

	SetDataDirectory("../data")
	feeds, err := RetrieveFeeds()

	if err != nil {
		t.Fatal(fmt.Sprintf("cannot load the %v: %v", dataFile, err), ballotX)
	}

	if len(feeds) == 0 {

		t.Fatal("no feed has been loaded!\n", ballotX)
	}

	t.Log("feeds have been loaded", checkMark)

}
