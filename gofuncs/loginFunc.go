//login related functions
package hallo

import (
	"net/http"
)

func retriveLoginData (r *http.Request) (*Params,error) {

    user := r.FormValue("username")
    pass := r.FormValue("password")
    
    if user != "" && pass != "" {
        param := Params {"Username" :  user,"Password" : pass}
        return &param,nil
    }
    return nil,nil
}

func checkIdentity (u , pass []string) bool {
    h := true

    return h
}


func processLoginFunc (r *http.Request,param *Params) error {
	return nil
}
