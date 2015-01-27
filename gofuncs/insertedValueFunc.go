//insertedValue functions
package hallo

import (
	"net/http"
	"time"
	"appengine"
    "appengine/datastore"
    "log"
)


// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}
func retriveInsertedValueData (r *http.Request) *Params {
    
    date := r.FormValue("date")

    if date != "" {
        param := Params {"Date" : date}
        return &param
    }
    return nil
}
func processInsertedValueData (r *http.Request,params *Params) error  {
        c := appengine.NewContext(r)
        value := (*Params)["Date"]

        str, ok := value.(string); 
    		
        if ok! { 

        }
		
		date := Date {
            Person: "WhoWasInserting",
            Content: value,
            Date:    time.Now(),

        log.Println(date)
        key := datastore.NewIncompleteKey(c, "Date", guestbookKey(c))
        _, err := datastore.Put(c, key, &date)
        return err
}

