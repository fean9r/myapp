package hallo

import (
	"net/http"
    "time"
    "fmt"
    "io"    
    "crypto/md5"
    "strconv"
    "log"
)


func dayFunc() *Params {
    tc := Params {"Day" : time.Now()}
    return &tc
}

func insertPageFunc() *Params {
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    tc := Params {"Token" : token}
    return &tc
}

func loginFunc() *Params {
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    tc := Params {"Token" : token}
    return &tc
}

func loadParams(title string) *Params {
    log.Println("load Params Titled", title)
    switch title {
        case "day":
            return dayFunc()
        case "insertPage":
            return insertPageFunc()
        case "login":
            return loginFunc()
        default:
            return nil
        }
}


func view(w http.ResponseWriter, r *http.Request,title string ) {

    params := loadParams(title)
    renderTemplate(w,title,params)

}



