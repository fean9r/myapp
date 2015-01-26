//login related functions
package hallo

import (
	"net/http"
)

func retriveLoginData (r *http.Request) *Params {

    user := r.FormValue("username")
    pass := r.FormValue("password")
    
    if user != "" && pass != "" {
        param := Params {"Username" :  user,"Password" : pass}
        return &param
    }
    return nil
}

func checkIdentity (u , pass []string) bool {
    h := true

    return h
}


func processLoginFunc (param *Params) {
	
}
