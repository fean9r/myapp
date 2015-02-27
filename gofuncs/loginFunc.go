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

// view handler func
func loginFunc(w http.ResponseWriter, r *http.Request) (*Params,error) {
    
    sess := globalSessions.SessionStart(w, r)
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    
    param := Params {
        "Token" : token,
        "username" : sess.Get("username")}
    return &param , nil
}

// insert handler func
func retriveLoginData (w http.ResponseWriter,r *http.Request) (*Params,error) {
    sess := globalSessions.SessionStart(w, r)
    user := r.FormValue("username")
    pass := r.FormValue("password")
    
    if goodIdentity(&user, &pass) {
        param := Params {
            "Username" :  user,
            "Password" : pass,
        }
        sess.Set("username", user)
        return &param,nil
    }
    return nil,errors.New("Not correct login data")
}

func goodIdentity (u , pass *string) bool {
    h := true

    if *u == "" || *pass=="" {
        h=false
    }

    return h
}


func processLoginFunc (w http.ResponseWriter, r *http.Request,param *Params) error {
    
	return nil
}


