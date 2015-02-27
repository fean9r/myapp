//day page related functions
package specificPages

import (
    "time"
    "appengine"
    "appengine/datastore"
)

func dayFunc(w http.ResponseWriter, r *http.Request) (*Params,error) {
    
    c := appengine.NewContext(r)
        // Ancestor queries, as shown here, are strongly consistent with the High
        // Replication Datastore. Queries that span entity groups are eventually
        // consistent. If we omitted the .Ancestor from this query there would be
        // a slight chance that Greeting that had just been written would not
        // show up in a query.
    q := datastore.NewQuery("Date").Ancestor(guestbookKey(c)).Filter("Person =", "WhoWasInserting").Order("-Date").Limit(10)

    dates := make([]Date, 0, 10)
    _, err := q.GetAll(c, &dates)

    param := Params {
        "Day" : time.Now(),
        "Dates" : dates,
        }
    return &param , err
}