//insertedValue functions
package hallo

import (
	"net/http"
	"time"
	"appengine"
    "appengine/datastore"
    "log"
    "errors"
    "io"    
    "crypto/md5"
    "fmt"
    "strconv"
)


// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

func checkDate ( value *string) bool {
    h := true

    if *value == ""  {
        h=false
    }

    // add regex here!
    return h
}

func insertPageFunc (w http.ResponseWriter,r *http.Request) (*Params,error) {
    
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    param := Params {"Token" : token}
    return &param , nil
}


func retriveInsertedValueData (w http.ResponseWriter,r *http.Request) (*Params,error) {
    
    if date := r.FormValue("date") ; checkDate(&date) {
        param := Params {"Date" : date}
        return &param ,nil
    }

    return nil,errors.New("Not correct date")
}
func processInsertedValueData (w http.ResponseWriter,r *http.Request,params *Params) error  {
        c := appengine.NewContext(r)
        value := (*params)["Date"]

        str, ok := value.(string); 
    		
        if ok!=true { 
        	return errors.New("String cast Exception")
        }
		
		date := Date {
            Person: "WhoWasInserting",
            Content: str,
            Date: time.Now(),
		}

        log.Println(date)
        key := datastore.NewIncompleteKey(c, "Date", guestbookKey(c))
        _, err := datastore.Put(c, key, &date)
        return err
}

