package hallo

import (
	"fmt"
	"html/template"
	"time"
	"net/http"
	"log"
	"crypto/md5"
	"io"
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

var global_var int

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/day", day)
	http.HandleFunc("/week", week)
	http.HandleFunc("/insertPage", insertPage)
	http.HandleFunc("/stats", stats)


}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, topForm)
	t, _ := template.ParseFiles("Main.gtpl")
    t.Execute(w, nil)
	fmt.Fprint(w, bottomForm)
}

const topForm = `
<html ng-app="schedulerApp">
<head >
	<meta charset="utf-8">
	<link rel="stylesheet" href="stylesheets/myapp.css">
	<link rel="stylesheet" href="lib/scheduler/dhtmlxscheduler.css">

	<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
	<script src="lib/angular/angular.min.js"></script>
	<script src="lib/scheduler/dhtmlxscheduler.js"></script>
	<script src="scripts/app.js"></script>
	<script src="scripts/app.scheduler.js"></script>
	<script src="scripts/myscripts.js"></script>
	<script type="text/javascript">
		// some javascript script 
	</script>

</head>
  <body class="app" ng-controller="MainSchedulerCtrl">
  	<div class="header">
  		<h1 class="logo " id="logo">
  			<a class="logo_anchor" href="/">MyApp</a>
  		</h1>
  		<a class="login" href="/login">Login</a>
  	</div> <!-- End header-->
  	<div class="container">
`
const bottomForm = `
	
		<footer class="site-footer">
    		<p class="upcomer">All rights reserved</p>
    	</footer>
	</div> <!-- End Container-->
  </body>
</html>
`

func day( w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, topForm)
	tim := time.Now()
	t, _ := template.ParseFiles("Day.gtpl")
    t.Execute(w, tim)
	fmt.Fprint(w, bottomForm)
}


func week( w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, topForm)
	t, _ := template.ParseFiles("Week.gtpl")
    t.Execute(w, nil)
	fmt.Fprint(w, bottomForm)
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}
func insertPage( w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, topForm)

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
	
	fmt.Fprint(w, bottomForm)
}

func stats( w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, topForm)
	t, _ := template.ParseFiles("stats.gtpl")
    err := t.Execute(w, nil)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, bottomForm)
}

func checkIdentity(u , pass []string) bool {
	h := true

	return h
}

func login( w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, topForm)
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
	fmt.Fprint(w, bottomForm)
}

