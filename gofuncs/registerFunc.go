//register related functions
package hallo

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
    "time"
    "appengine"
    "appengine/datastore"
	)

// userbookKey returns the key used for all Userbook entries.
func userbookKey(c appengine.Context) *datastore.Key {
    // The string "default_Userbook" here could be varied to have multiple Userbooks.
    return datastore.NewKey(c, "Userbook", "default_Userbook", 0, nil)
}


// internal function to check if the inserted values are ok is good and known in the system
func goodRegister (u , pass *string) bool {
    h := true

    if *u == "" || *pass=="" {
        h=false
    }

    return h
}

func registerPageFunc (w http.ResponseWriter, r *http.Request, param *Params) (error) {
    
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    (*param)["Token"] = token
    return nil
}

// function that retreives the data from the page and makes an internal verification
// to insure correct insertions
func retriveInsertedRegisterData (w http.ResponseWriter,r *http.Request,param *Params) (error) {
	
	user, pass := r.FormValue("username"),r.FormValue("password")
	if goodRegister(&user, &pass) {
        (*param)["Username"] = user
        (*param)["Password"] = pass
        return nil
    }
	return errors.New("Incorrect Register data")
}

// function that uses the inserted data that are now verified 
func processRegisterData (w http.ResponseWriter, r *http.Request, params *Params) (error) {
    
    username,pass := (*params)["Username"] ,(*params)["Password"] 
    strUsername, ok := username.(string);       
    if ok!=true { 
        return errors.New("String cast Exception")
    }
    strPass, ok := pass.(string);       
    if ok!=true { 
        return errors.New("String cast Exception")
    }

    user := User {
		Username: strUsername,
	}
	user.SetPassword(strPass)

	c := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(c, "User", userbookKey(c))
    _, err := datastore.Put(c, key, &user)
	//reportValue("processRegisterData ",user)
	return err
}