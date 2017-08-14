package search

// The interface to be implemented for any supported Matcher,
// it will be the extension point of this application
type Matcher interface{
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// The structure of the expected results from matcher's searching
type Result struct {
	Field string
	Content string
}