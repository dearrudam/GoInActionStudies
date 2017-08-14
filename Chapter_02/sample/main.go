package main

import (
	"log"
	"os"
	_ "github.com/dearrudam/GoInActionStudies/Chapter_02/sample/matchers"
	"github.com/dearrudam/GoInActionStudies/Chapter_02/sample/search"
)

//init is called prior to main
func init() {
	// set the device for logging  to stdout
	log.SetOutput(os.Stdout)
}

//main is the entry point for the program
func main() {
	//perform the search for the specific term
	search.Run("president")
}
