package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
    "os"
    "path/filepath"
    "strconv"
	"regexp"
)


//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and 
// if responsewriter is passed, outputs to http 

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
    return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")

        for _, r := range regexes {
            if r.MatchString(path) {
                var tfile FileInfo
                dir, filename := filepath.Split(path)
                tfile.Filename = string(filename)
                tfile.Location = string(dir)

                //TODO_5: As it currently stands the same file can be added to the array more than once 
                //TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
                Files = append(Files, tfile)

                if w != nil && len(Files)>0 {

                    //TODO_6: The current key value is the LEN of Files (this terrible); 
                    //TODO_6: Create some variable to track how many files have been added
                    w.Write([]byte(`"`+(strconv.FormatInt(int64(len(Files)), 10))+`":  `))
                    json.NewEncoder(w).Encode(tfile)
                    w.Write([]byte(`,`))

                } 
                
                log.Printf("[+] HIT: %s\n", path)

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
	return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")

        r := regexp.MustCompile(`(?i)+ query`)

    	if r.MatchString(path) {
                var tfile FileInfo
                dir, filename := filepath.Split(path)
                tfile.Filename = string(filename)
                tfile.Location = string(dir)

                //TODO_5: As it currently stands the same file can be added to the array more than once 
                //TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record

                Files = append(Files, tfile)

                if w != nil && len(Files)>0 {

                    //TODO_6: The current key value is the LEN of Files (this terrible); 
                    //TODO_6: Create some variable to track how many files have been added
                    w.Write([]byte(`"`+(strconv.FormatInt(int64(len(Files)), 10))+`":  `))
                    json.NewEncoder(w).Encode(tfile)
                    w.Write([]byte(`,`))

                } 
                
                log.Printf("[+] HIT: %s\n", path)

        }
        return nil
    }
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {

	log.Printf("Entering %s end point", r.URL.Path)
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
	log.Println(regexes)

}


func MainPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
    w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	
    //TODO_8 - Write out something better than this that describes what this api does
	//w.Write([]byte(`{ "status" : "Welcome to the Main Page ",`))
	fmt.Fprintf(w, "<html><body>API with features to inlude search,indexer,addsearch,clear and reset<H1></H1></body>")

}


func FindFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
    q, ok := r.URL.Query()["q"]

    w.WriteHeader(http.StatusOK)
    if ok && len(q[0]) > 0 {
        log.Printf("Entering search with query=%s",q[0])

        // ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
        // e.g., func finder(query string) []string { ... }
			//FOUND := false
        for _, File := range Files {
		    if File.Filename == q[0] {
                json.NewEncoder(w).Encode(File.Location)
                //FOUND := true
		    }else{
				w.Write([]byte(`{ "Status" : "FILE NOT FOUND ",`))
			}
        }
        //TODO_9: Handle when no matches exist; print a useful json response to the user; hint you might need a "FOUND variable" to check here ...

    } else {
        // didn't pass in a search term, show all that you've found
        w.Write([]byte(`"files":`))    
        json.NewEncoder(w).Encode(Files)
    }
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
    log.Printf("Entering %s end point", r.URL.Path)
    w.Header().Set("Content-Type", "application/json")

    location, locOK := r.URL.Query()["location"]
	regex, regexOK := r.URL.Query()["regex"]
    //TODO_10: Currently there is a huge risk with this code ... namely, we can search from the root /
    //TODO_10: Assume the location passed starts at /home/ (or in Windows pick some "safe?" location)
    //TODO_10: something like ...  rootDir string := "???"
    //TODO_10: create another variable and append location[0] to rootDir (where appropriate) to patch this hole
	rootDir := "/home/cabox"
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

    //wrapper to make "nice json"
    w.Write([]byte(`{ `))
    
    // TODO_11: Currently the code DOES NOT do anything with an optionally passed regex parameter
    // Define the logic required here to call the new function walkFn2(w,regex[0])
    // Hint, you need to grab the regex parameter (see how it's done for location above...) 
    
    if regexOK{
		filepath.Walk(rootDir + location[0], walkFn2(w, regex[0]))
	}else {
		filepath.Walk(rootDir + location[0], walkFn(w))
	}
    //   call 
    // else run code to locate files matching stored regular expression
 

    //wrapper to make "nice json"
    w.Write([]byte(` "status": "completed"} `))

}


//TODO_12 create endpoint that calls resetRegEx AND *** clears the current Files found; ***
//TODO_12 Make sure to connect the name of your function back to the reset endpoint main.go!

func RESETARRAY(w http.ResponseWriter, r *http.Request) {
    log.Printf("Entering %s end point", r.URL.Path)
    w.Header().Set("Content-Type", "application/json")
	resetRegEx()

}
//TODO_13 create endpoint that calls clearRegEx ; 
func CLEARARRAY(w http.ResponseWriter, r *http.Request) {
    log.Printf("Entering %s end point", r.URL.Path)
    w.Header().Set("Content-Type", "application/json")
	clearRegEx()
}
//TODO_12 Make sure to connect the name of your function back to the clear endpoint main.go!


//TODO_14 create endpoint that calls addRegEx ;

func AddRegEx(w http.ResponseWriter, r *http.Request) {
    log.Printf("Entering %s end point", r.URL.Path)
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	addRegEx(`(?i)` + params["regex"])
}
//TODO_12 Make sure to connect the name of your function back to the addsearch endpoint in main.go!
// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint
