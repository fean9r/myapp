//insertedValue functions
package hallo

import (
	"net/http"
	"appengine"
    "appengine/datastore"
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
func processInsertedValueData (params *Params)  {
        // c := appengine.NewContext(r)
        // date := Date{
        //         Person: "mao",
        //         Content: r.FormValue("date"),
        //         Date:    time.Now(),
        // }
        // key := datastore.NewIncompleteKey(c, "Date", guestbookKey(c))
        // _, err := datastore.Put(c, key, &date)
        // if err != nil {
        //         http.Error(w, err.Error(), http.StatusInternalServerError)
        //         return
        // }
}