package search


type Matcher interface{
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

type Result struct {
	Field string
	Content string
}