package hallo

import (
	"net/http"
)


func getParams (title string,r *http.Request) *Params {
    switch title {
        case "insertedValue":
            return retriveInsertedValueData(r)
        case "login":
            return retriveLoginData(r)
        default:
            return nil
        }
}


func processRequest (title string,param *Params)  {
    
    switch title {
        case "insertedValue":
            processInsertedValueData(param)
        case "login":
            processLoginFunc(param)
        default:
             
        }
}


func insert(w http.ResponseWriter, r *http.Request, title string) {

    if r.Method == "GET" {
        http.Error(w, "NOT POSSIBLE", http.StatusInternalServerError)
    } else {
       
        r.ParseForm()
        
        token := r.Form.Get("token")
        
        param := getParams(title,r)
        if  token != "" && param != nil  {
            processRequest(title,param)
            renderTemplate(w,title,param)
        } else {
            // error 
            // duplicate submission 
            // empty value
            renderTemplate(w,title+"Error",nil)
        }
    }
}