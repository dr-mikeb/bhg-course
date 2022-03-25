package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"github.com/gorilla/mux"
	
)

//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and
// if responsewriter is passed, outputs to http

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
	indexTracker := 0
    return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")
		
        for _, r := range regexes {
            if r.MatchString(path) {
                var tfile FileInfo
                dir, filename := filepath.Split(path)//seperates by directory and filename
                tfile.Filename = string(filename)//store in data structure
                tfile.Location = string(dir)//store in data structure

                //TODO_5: As it currently stands the same file can be added to the array more than once 
                //TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
				//TODO_5:ISAIAH check for duplicates to prevent duplicate entries
				//make sure it doesn't already exist before you add it
				var duplicate bool
				duplicate = false
				for _, f := range Files {
					if(f.Filename == tfile.Filename) &&(f.Location == tfile.Location){
						duplicate = true
					}
				}
				if !duplicate {
					Files = append(Files, tfile)//append to data structrue

                	if w != nil && len(Files)>0 {

                    	//TODO_6: The current key value is the LEN of Files (this terrible); 
                    	//TODO_6: Create some variable to track how many files have been added
						//its currently using the length of the file as the index, come up with something better
						//keep track of an index
                    	//w.Write([]byte(`"`+(strconv.FormatInt(int64(len(Files)), 10))+`":  `))
						w.Write([]byte(`"`+(strconv.FormatInt(int64(indexTracker), 10))+`":  `))
						indexTracker++
                    	json.NewEncoder(w).Encode(tfile)
                    	w.Write([]byte(`,`))

                	} 
					if LOG_LEVEL==2{
                		log.Printf("[+] HIT: %s\n", path)
					}
				} else{
					if(LOG_LEVEL==2){
						log.Printf("Duplicate file!\n")
					}
				}
            }

        }
        return nil
    }

}

//TODO_7: One of the options for the API is a query command
//TODO_7: Create a walkFn2 function based on the walkFn function, 
//TODO_7: Instead of using the regexes array, define a single regex 
//TODO_7: Hint look at the logic in scrape.go to see how to do that; 
//TODO_7: You won't have to itterate through the regexes for loop in this func!

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
    indexTracker := 0
    return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")
		
        // for _, r := range regexes {
		r := regexp.MustCompile(`(?i)`+query)
		if r.MatchString(path) {
			var tfile FileInfo
			dir, filename := filepath.Split(path)//seperates by directory and filename
			tfile.Filename = string(filename)//store in data structure
			tfile.Location = string(dir)//store in data structure

			var duplicate bool
			duplicate = false
			for _, f := range Files {
				if(f.Filename == tfile.Filename) &&(f.Location == tfile.Location){
					duplicate = true					}
			}
			if !duplicate {
				Files = append(Files, tfile)//append to data structrue

			if w != nil && len(Files)>0 {

					//TODO_6: The current key value is the LEN of Files (this terrible); 
					//TODO_6: Create some variable to track how many files have been added
					//its currently using the length of the file as the index, come up with something better
					//keep track of an index
					//w.Write([]byte(`"`+(strconv.FormatInt(int64(len(Files)), 10))+`":  `))
					w.Write([]byte(`"`+(strconv.FormatInt(int64(indexTracker), 10))+`":  `))
					indexTracker++
					json.NewEncoder(w).Encode(tfile)
					w.Write([]byte(`,`))

				} 
                if(LOG_LEVEL==2){
					log.Printf("[+] HIT: %s\n", path)
				}
			} else{
				if(LOG_LEVEL==2){
					log.Printf("Duplicate file!\n")
				}
			}
			
		}

        return nil
    }
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL==1 || LOG_LEVEL==2{
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{ "status" : "API is up and running ",`))
    var regexstrings []string
    
    for _, regex := range regexes{
        regexstrings = append(regexstrings, regex.String())
    }

    w.Write([]byte(` "regexs" :`))
    json.NewEncoder(w).Encode(regexstrings)
    w.Write([]byte(`}`))
	if LOG_LEVEL == 2{
		log.Println(regexes)
	}
}


func MainPage(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL==1 || LOG_LEVEL==2{
		log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
    //TODO_8 - Write out something better than this that describes what this api does
	//Isaiah just modify html
	fmt.Fprintf(w, "<html><body><H1>This is the Main Page:</H1><h2>You can use this to index, edit, and search your filesystem!</h2></body>")
}

//Isaiah used for search
func FindFile(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL==1 || LOG_LEVEL==2{
		log.Printf("Entering %s end point", r.URL.Path)
	}
    q, ok := r.URL.Query()["q"]

    w.WriteHeader(http.StatusOK)
    if ok && len(q[0]) > 0 {
		if LOG_LEVEL==1 || LOG_LEVEL==2{
        	log.Printf("Entering search with query=%s",q[0])
		}

        // ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
        // e.g., func finder(query string) []string { ... }
		found := false
        for _, File := range Files {
		    if File.Filename == q[0] {
                json.NewEncoder(w).Encode(File.Location)
                //consider FOUND = TRUE
				found = true
		    }
        }
        //TODO_9: Handle when no matches exist; print a useful json response to the user; hint you might need a "FOUND variable" to check here ...
		if(!found){
			w.Write([]byte(`"No files with this search query could be found"`)) 
		 }
    } else {
        // didn't pass in a search term, show all that you've found
        w.Write([]byte(`"files":`))    
        json.NewEncoder(w).Encode(Files)
    }
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL==1 || LOG_LEVEL==2{
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")

    location, locOK := r.URL.Query()["location"]
    
    //TODO_10: Currently there is a huge risk with this code ... namely, we can search from the root /
    //TODO_10: Assume the location passed starts at /home/ (or in Windows pick some "safe?" location)
    //TODO_10: something like ...  rootDir string := "???"
    //TODO_10: create another variable and append location[0] to rootDir (where appropriate) to patch this hole
	rootDir := `/home/cabox`
	//append to location[0] after error checks
    if locOK && len(location[0]) > 0 {
        w.WriteHeader(http.StatusOK)

    } else {
        w.WriteHeader(http.StatusFailedDependency)
        w.Write([]byte(`{ "parameters" : {"required": "location",`))    
        w.Write([]byte(`"optional": "regex"},`))    
        w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
        w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
        return 
    }
	location[0] = rootDir+location[0]
    //wrapper to make "nice json"
    w.Write([]byte(`{ `))
    
    // TODO_11: Currently the code DOES NOT do anything with an optionally passed regex parameter
    // Define the logic required here to call the new function walkFn2(w,regex[0])
    // Hint, you need to grab the regex parameter (see how it's done for location above...) 
    //ISaiah related to todo7
	regex, regexOK := r.URL.Query()["regex"]
	
    if regexOK {
		filepath.Walk(location[0], walkFn2(w, regex[0]))
	}else{
		if err := filepath.Walk(location[0], walkFn(w)); err != nil {
			log.Panicln(err)
		}
	}
    //   call filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
    // else run code to locate files matching stored regular expression
    // if err := filepath.Walk(location[0], walkFn(w)); err != nil {
	// 	log.Panicln(err)
	// }
    //wrapper to make "nice json"
    w.Write([]byte(` "status": "completed"} `))

}


//TODO_12 create endpoint that calls resetRegEx AND *** clears the current Files found; ***
//Files = nil
//TODO_12 Make sure to connect the name of your function back to the reset endpoint main.go!

func ResetArray(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL==1 || LOG_LEVEL==2{
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")
	resetRegEx()
	Files = nil
	w.Write([]byte(`{ "status": "completed"} `))
}

//TODO_13 create endpoint that calls clearRegEx ; 
//TODO_13 Make sure to connect the name of your function back to the clear endpoint main.go!
func ClearArray(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL==1 || LOG_LEVEL==2{
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")
	clearRegEx()
	w.Write([]byte(`{ "status": "completed"} `))
}


//TODO_14 create endpoint that calls addRegEx ; 
//TODO_14 Make sure to connect the name of your function back to the addsearch endpoint in main.go!
// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint


func AddRegExpression(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL==1 || LOG_LEVEL==2{
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if len(params["regex"]) > 0 {
		addRegEx(`(?i)`+params["regex"])
	}else{
		if LOG_LEVEL==1 || LOG_LEVEL==2{
			log.Printf("No Expression Entered\n")
		}
	}
	w.Write([]byte(`{ "status": "completed"} `))
}
