package hallo

import (
	"net/http"
    "time"
    "fmt"
    "io"    
    "crypto/md5"
    "strconv"
    "log"
   // "errors"
    "appengine"
    "appengine/datastore"
)


func dayFunc(r *http.Request) (*Params,error) {
    
    c := appengine.NewContext(r)
        // Ancestor queries, as shown here, are strongly consistent with the High
        // Replication Datastore. Queries that span entity groups are eventually
        // consistent. If we omitted the .Ancestor from this query there would be
        // a slight chance that Greeting that had just been written would not
        // show up in a query.
    q := datastore.NewQuery("Date").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
    dates := make([]Date, 0, 10)
    _, err := q.GetAll(c, &dates)

    tc := Params {
        "Day" : time.Now(),
        "Dates" : dates,
        }
    return &tc , err
}

func insertPageFunc(r *http.Request) (*Params,error) {
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    tc := Params {"Token" : token}
    return &tc , nil
}

func loginFunc(r *http.Request) (*Params,error) {
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    tc := Params {"Token" : token}
    return &tc , nil
}


type fn func (*http.Request) (*Params,error)


func loadParams(r *http.Request,title string) (*Params, error) {
    log.Println("load Params Titled", title)

    m := map[string] fn {
        "day": dayFunc,
        "insertPage": insertPageFunc,
        "login": loginFunc,
        "week": nil,
        "stats": nil,
        "main": nil,
    }
    return m[title](r)

    // switch title {
    //     case "day":
    //         return dayFunc(r)
    //     case "insertPage":
    //         return insertPageFunc()
    //     case "login":
    //         return loginFunc()
    //     case "week":
    //         return nil,nil
    //     case "stats":
    //         return nil,nil
    //     case "main":
    //         return nil,nil
    //     default:
    //         return nil,errors.New(title+" handling function not found")
    //     }
}


func view(w http.ResponseWriter, r *http.Request,title string ) {

    params,err := loadParams(r,title)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    if params!=nil {
        log.Println(*params)
    }
    renderTemplate(w,title,params)

}



