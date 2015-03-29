//insertedValue functions
package hallo

import (
	"net/http"
	"time"
	"appengine"
    "appengine/datastore"
    "errors"
    "io"    
    "crypto/md5"
    "fmt"
    "strconv"
)


// datebookKey returns the key used for all Datebook entries.
func datebookKey(c appengine.Context) *datastore.Key {
    // The string "default_Datebook" here could be varied to have multiple Datebooks.
    return datastore.NewKey(c, "Datebook", "default_Datebook", 0, nil)
}



func checkDate ( value *string) bool {
    h := true

    if *value == ""  {
        h=false
    }

    // add regex here!
    return h
}


func insertPageFunc (w http.ResponseWriter,r *http.Request,param *Params) (error) {
    
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    (*param)["Token"] = token
    return nil
}


func retriveInsertedValueData (w http.ResponseWriter,r *http.Request, param *Params) (error) {
    
    if date := r.FormValue("date") ; checkDate(&date) {
         (*param)["Date"] = date
        return nil
    }
    return errors.New("Not correct date")
}

func processInsertedValueData (w http.ResponseWriter,r *http.Request,params *Params) error  {
        
        value := (*params)["Date"]

        str, ok := value.(string); 
    		
        if ok!=true { 
        	return errors.New("String cast Exception")
        }
		
		date := Date {
            User: "WhoWasInserting",
            Content: str,
            Date: time.Now(),
		}

        reportValue("processInsertedValueData",date)
        c := appengine.NewContext(r)
        key := datastore.NewIncompleteKey(c, "Date", datebookKey(c))
        _, err := datastore.Put(c, key, &date)
        return err
}

