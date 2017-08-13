package search

import (
	"log"
	"sync"
	"fmt"
)

var matchers = make(map[string]Matcher)

func Register(name string, matcher Matcher) {
	matchers[name] = matcher
}

func Run(searchTerm string) {

	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {

		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	returnedResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Fatalf("failure on search \"%v\" on feed \"%v\" : %v", searchTerm, feed.Name, err)
	}
	for _, result := range returnedResults {
		results <- result
	}
}

func Display(results chan*Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
