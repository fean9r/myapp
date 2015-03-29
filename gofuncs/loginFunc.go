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
    "appengine"
    "appengine/datastore"
    "code.google.com/p/go.crypto/bcrypt"
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

//Login validates and returns a user object if they exist in the database.
func Login(r *http.Request,ctx *Params, username, password string) (u *User, err error) {

    c := appengine.NewContext(r)
    q := datastore.NewQuery("User").Ancestor(userbookKey(c)).Filter("Username =", username).Limit(1)
    users := make([]User, 0, 2)
    _, err = q.GetAll(c, &users)
    //err = ctx.C("users").Find(bson.M{"username": username}).One(&u)
    if err != nil {
        return nil, err
    }
    // query done without problems
    if len(users) < 1 {
        //return nil, errors.New("Incorrect Username or password.")
        return nil, errors.New("Username not found!")
    }
    //retreives something, now check

    u = &User { 
        Username : users[0].Username,
        Password: users[0].Password,
    }
    reportValue("Login",u)
    err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
    
    if err != nil {
        return nil, errors.New("Password not found!")
    }
    return
}


//  the xxxPageFunc will provide the data for the view the page 
func loginPageFunc (w http.ResponseWriter, r *http.Request, param *Params) (error) {
    
    sess := globalSessions.SessionStart(w, r)
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    (*param)["Token"] = token
    reportValue("loginPageFunc" , sess.Get("username"))
    return nil
}

// function that retreives the data from the page and makes an internal verification
// to insure correct insertions
func retriveInsertedLoginData (w http.ResponseWriter, r *http.Request, param *Params) (error) {
    
    user, pass := r.FormValue("username"), r.FormValue("password")
    if goodIdentity(&user, &pass) {
        (*param)["Username"] = user
        (*param)["Password"] = pass
        return nil
    }
    return errors.New("Not correct login data")
}



func processLoginData (w http.ResponseWriter, r *http.Request, params *Params) (error) {
    username, pass := (*params)["Username"], (*params)["Password"]  

    strUsername, ok := username.(string);       
    if ok!=true { 
        return errors.New("String cast Exception")
    }
    strPass, ok := pass.(string);       
    if ok!=true { 
        return errors.New("String cast Exception")
    }

    user, err := Login(r,params,strUsername,strPass)
    
    if err != nil {
        return err
    }

    (*params)["User"] = user
    //sess := globalSessions.SessionStart(w, r)
    //sess.Set("username", user)
    //log.Println("SessionLog ",sess.Get("username"))
	return nil
}


