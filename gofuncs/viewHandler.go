package hallo

import (
	"net/http"
)

func view(w http.ResponseWriter, r *http.Request, title string ) {

    viewPageFunc,err := viewFuncManager.getFunction(title)
    if err != nil {
        renderTemplate(w,"Error",err)
        return
    }

    if viewPageFunc == nil {
        renderTemplate(w,title,nil)
    }else{

        param := Params {}
        err := viewPageFunc(w,r,&param)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        // not usefull just for testing
        if param != nil {
            reportValue("view",param)
        }
        renderTemplate(w,title,param)
    }
}



