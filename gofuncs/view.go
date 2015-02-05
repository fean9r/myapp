package hallo

import (
	"net/http"
    "time"
    "fmt"
    "io"    
    "crypto/md5"
    "strconv"
    "log"
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

    param := Params {
        "Day" : time.Now(),
        "Dates" : dates,
        }
    return &param , err
}

func insertPageFunc(r *http.Request) (*Params,error) {
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    param := Params {"Token" : token}
    return &param , nil
}

func loginFunc(r *http.Request) (*Params,error) {
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    param := Params {"Token" : token}
    return &param , nil
}


func view(w http.ResponseWriter, r *http.Request,title string ) {

    //params,err := loadParams(r,title)
    myfunc,err := viewFuncManager.getFunction(title)
    if err != nil {
        renderTemplate(w,title,nil)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
    }else {
        params,err := myfunc(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        if params!=nil {
            log.Println(*params)
        }
        renderTemplate(w,title,params)
    }
}



