package hallo

import (
	"html/template"
	"net/http"
	"regexp"
	//"log"
)

var HandleFunctions = ""

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

var templates = template.Must(template.ParseGlob("gtpl/*"))

type Params map[string]interface {}


func renderTemplate (w http.ResponseWriter, tmpl string, p *Params) {

    // if p != nil {
    //     log.Println("Template params", *p)
    // }
    err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func init() {
	
	addHandler("/insert/", makeHandler(insert))
    addHandler("/view/", makeHandler(view))
    validPath = regexp.MustCompile("^/("+ HandleFunctions +")/([a-zA-Z0-9]+)$")
}

