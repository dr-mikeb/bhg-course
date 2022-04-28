package scrape

// scrape.go HAS ONE TODO - TODO_15 and one OPTIONAL "ADVANCED" ASK and one OPTIONAL CHALLENGE
// the advanced ask has some (minimal) implications to scrapeapi.go while the advanced ask has substantial implications


import (
	"regexp"
)


//==========================================================================\\
// || GLOBAL DATA STRUCTURES  ||

//ADVANCED: This is perhaps a terrible structure since multiple a filename is NOT guarenteed to be unique; consider an array of Locations? 
//CHALLENGE: Replace this Local Structure with a Key-Value DB like REDIS
type FileInfo struct {
	Filename string `json:"filename"`
	Location string `json:"location"`
}
var Files []FileInfo



var regexes = []*regexp.Regexp{
	regexp.MustCompile(`(?i)password`),
    regexp.MustCompile(`(?i).txt`),  
}

// END GLOBAL VARIABLES
//==========================================================================//

//==========================================================================\\
// || HELPER FUNCTIONS TO MANIPULATE THE REGULAR EXPRESSIONS ||

func resetRegEx(){
    regexes = []*regexp.Regexp{
        regexp.MustCompile(`(?i)password`),
        regexp.MustCompile(`(?i)kdb`),
        regexp.MustCompile(`(?i)login`),
    }
}

func clearRegEx(){
     //TODO_15 - Validate that this works as expected and doesn't cause issues
    regexes = nil
}

func addRegEx(regex string){
    regexes = append(regexes,regexp.MustCompile(regex))
}

//==========================================================================//





