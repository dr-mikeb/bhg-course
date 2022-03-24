# Lab 6
Due April 1st at 11:59PM

## Development Work [10 points; ~2/3 point per todo]
- Complete the 4 TODOs in [main.go](course-materials/materials/lab/6/main/main.go)
- Complete the 1 TODO in [scrape/scrape.go](course-materials/materials/lab/6/scrape/scrape.go)
- Complete the 10 TODOs in [scrape/scrapeapi.go](course-materials/materials/lab/6/scrape/scrapeapi.go)


## Capture Submission details

- Save the html files containing the response to the following queries to your API; 
  - note in some cases you need to run mulitple API calls prior to capturing the html output 
- Name the files as directed
- Replace YOURHOST below as appropriate; Replace LASTNAME as appropriate, execute the API calls on a fresh server in the following order


0. http://YOURHOST:8080/indexer?location=/  --> [LASTNAME_ROOTCHECK.HTML]
2. http://YOURHOST:8080/clear  ;  http://YOURHOST:8080/addsearch/go ; http://YOURHOST:8080/api-status  --> [LASTNAME_CLEARSETCHECK.HTML]
3. http://YOURHOST:8080/search?q=main.md --> [LASTNAME_CHECKSEARCH.HTML]
4. http://YOURHOST:8080/search  --> [LASTNAME_ALL_PRERESET.HTML]
5. http://YOURHOST:8080/reset  ; http://YOURHOST:8080/indexer?location=/&regex=md --> [LASTNAME_CUSTOMREGEX.HTML]
6. http://YOURHOST:8080/search --> [LASTNAME_FINAL.HTML]

## Submit 
1. ZIP FILE with all 6 Files [10 points]
2. Link to your Github Repository [10 points]
3. List of Collaborators
