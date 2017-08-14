package search

// the default structure of Matcher
type defaultMatcher struct {
}

// making the defaultMatcher struct to implements the Matcher interface
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// init function is required because
// it will register the defaultMatcher on to
// supported matchers by the application
func init() {
	var matcher defaultMatcher
	// register the default Matcher into map of Matchers in to seach.go
	Register("default", matcher)
}

