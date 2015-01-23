package hallo

import (
	"fmt"
	"html/template"
	"time"
	"net/http"
	"log"
	"crypto/md5"
	"io"
    "io/ioutil"
	"strconv"
	"appengine"
	"appengine/datastore"
)

type Date struct {
        Person  string
        Content string
        Date    time.Time
}

type Person struct {
    Name string
    Password string
}

type Page struct {
    Title string
    Content  []byte
}
var global_var int

func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    content, err := ioutil.ReadFile("html/"+filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Content: content}, nil
}


var templates = template.Must(template.ParseGlob("gtpl/*"))

func init() {
   // http.HandleFunc("/", root)
	//http.HandleFunc("/login/", login)
	//http.HandleFunc("/day/", day)
	//http.HandleFunc("/week/", week)
	//http.HandleFunc("/insertPage/", insertPage)
	//http.HandleFunc("/stats/", stats)
    http.HandleFunc("/view/", view)

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
   // t, err := template.ParseFiles(tmpl + ".gtpl")
    //if err != nil {
    //    http.Error(w, err.Error(), http.StatusInternalServerError)
   // }
    err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func root(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, topForm)
    title := r.URL.Path[len("/???/"):]
    p, err := loadPage(title)
     if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        http.Redirect(w, r, "/???/"+title, http.StatusFound)
        return
    }
    // t, _ := template.ParseFiles("Main.gtpl")
    // t.Execute(w, nil)
    renderTemplate(w,"???",p)
	//fmt.Fprint(w, bottomForm)
}


func view(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprint(w, topForm)
    title := r.URL.Path[len("/view/"):]
    // p, err := loadPage(title)
    //  if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     http.Redirect(w, r, "/base/"+title, http.StatusFound)
    //     return
    // }
    // t, _ := template.ParseFiles("Main.gtpl")
    // t.Execute(w, nil)
   // renderTemplate(w,title,p)
    renderTemplate(w,title,nil)
    //fmt.Fprint(w, bottomForm)
}

func day( w http.ResponseWriter, r *http.Request) {
	
	tim := time.Now()
	t, _ := template.ParseFiles("Day.gtpl")
    t.Execute(w, tim)
}


func week( w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Week.gtpl")
    t.Execute(w, nil)
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}
func insertPage( w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
    	// prevent duplicate submission
   		crutime := time.Now().Unix()
    	h := md5.New()
    	io.WriteString(h, strconv.FormatInt(crutime, 10))
    	token := fmt.Sprintf("%x", h.Sum(nil))
   
   		// main page
		t, _ := template.ParseFiles("insertDate.gtpl")
    	t.Execute(w, token)
	} else {
		c := appengine.NewContext(r)
		r.ParseForm()
		
		date := Date{
				Person: "mao",
                Content: r.FormValue("date"),
                Date:    time.Now(),
        }
        // logic part of log in
        token := r.Form.Get("token")
        key := datastore.NewIncompleteKey(c, "Date", guestbookKey(c))
        _, err := datastore.Put(c, key, &date)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        val := r.FormValue("date")
        if  token != "" && val!= ""  {
			t, _ := template.ParseFiles("insertedValue.gtpl")
    		err := t.Execute(w, r.FormValue("date"))
    		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
        } else {
        	// error 
        	// duplicate submission 
       		// empty value
       	}
	}
	
}

func stats( w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("stats.gtpl")
    err := t.Execute(w, nil)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkIdentity(u , pass []string) bool {
	h := true

	return h
}

func login( w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method) //get request method
    if r.Method == "GET" {

    	//prevent  duplicate submission
    	crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))
        
        //
		t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, token)
    } else {
        r.ParseForm()
        // logic part of log in
        token := r.Form.Get("token")
        if checkIdentity(r.Form["username"],r.Form["password"]) && token != ""{
        	// happens when username and password are correct
        	//log.Println("username:", r.Form["username"])
        	//log.Println("password:", r.Form["password"])
        	t, _ := template.ParseFiles("loginSuccess.gtpl")
        	t.Execute(w, r.FormValue("username"))
    	} else {
    		// error
    		t, _ := template.ParseFiles("loginError.gtpl")
        	t.Execute(w, r.FormValue("username"))
    	}
    }
}

