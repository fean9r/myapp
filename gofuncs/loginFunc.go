//login related functions
package hallo

import (
	"net/http"
    "errors"
    "time"
    "crypto/md5"
    "fmt"
    "io"   
    "strconv"
)


// func login(w http.ResponseWriter, r *http.Request) {
//     sess := globalSessions.SessionStart(w, r)
//     r.ParseForm()
//     if r.Method == "GET" {
//         t, _ := template.ParseFiles("login.gtpl")
//         w.Header().Set("Content-Type", "text/html")
//         t.Execute(w, sess.Get("username"))
//     } else {
//         sess.Set("username", r.Form["username"])
//         http.Redirect(w, r, "/", 302)
//     }
// }

// internal function to check if the identity is good and known in the system
func goodIdentity (u , pass *string) bool {
    h := true

    if *u == "" || *pass=="" {
        h=false
    }

    return h
}


//  the xxxPageFunc will provide the data for the view the page 
func loginPageFunc (w http.ResponseWriter, r *http.Request, param *Params) (error) {
    
    //sess := globalSessions.SessionStart(w, r)
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    (*param)["Token"] = token
    //(*param)["username"] = sess.Get("username")
    return nil
}

// insert handler func
func retriveInsertedLoginData (w http.ResponseWriter,r *http.Request,param *Params) error {
    //sess := globalSessions.SessionStart(w, r)
    user := r.FormValue("username")
    pass := r.FormValue("password")
    
    if goodIdentity(&user, &pass) {
        (*param)["Username"] = user
        (*param)["Password"] = pass
        //sess.Set("username", user)
        return nil
    }
    return errors.New("Not correct login data")
}



func processLoginData (w http.ResponseWriter, r *http.Request, param *Params) (error) {
    
	return nil
}


