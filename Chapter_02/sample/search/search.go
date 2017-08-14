package search

import (
	"log"
	"sync"
	"fmt"
)

// A map of registered matchers for searching
var matchers = make(map[string]Matcher)

// A function responsible to register a given matcher with a given name
func Register(name string, matcher Matcher) {
	matchers[name] = matcher
}

// Run performs the search logic
func Run(searchTerm string) {

	// Retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create a unbuffered channel to receive match results
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {

		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done
	go func() {
		// Wait for everything to be processed
		waitGroup.Wait()
		// Close the channel to signal to the Display
		// function that we can exit the program
		close(results)
	}()

	// Start displaying results as they are available and
	// result after the final result is displayed
	Display(results)
}

// Perform the search and send all results to the results channel
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {

	// Searching for matched data based on provided term from the provided feed
	returnedResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Fatalf("failure on search \"%v\" on feed \"%v\" : %v", searchTerm, feed.Name, err)
	}
	// sending the received data to the results channel
	for _, result := range returnedResults {
		results <- result
	}
}

// Display the results on the stdout
func Display(results chan*Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
