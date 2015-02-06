//login related functions
package hallo

import (
	"net/http"
    "errors"
)

func retriveLoginData (r *http.Request) (*Params,error) {

    user := r.FormValue("username")
    pass := r.FormValue("password")
    
    if goodIdentity(&user, &pass) {
        param := Params {"Username" :  user,"Password" : pass}
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


func processLoginFunc (r *http.Request,param *Params) error {
    
	return nil
}
