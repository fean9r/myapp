package hallo

import (
	"net/http"
	"regexp"
)
var HandleFunctions string
var validPath  *regexp.Regexp

func makeHandler(fn func(http.ResponseWriter, *http.Request,string)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
        	http.NotFound(w, r)
       		return 
    	}
    	fn(w, r, m[2])
	}
}

func addHandler (title string, handler http.HandlerFunc ) {

	HandleFunctions+= title[1:len(title)-1]+"|"
	http.HandleFunc(title, handler)
}