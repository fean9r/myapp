//register related functions
package hallo

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
    "time"
	)

func registerPageFunc (w http.ResponseWriter, r *http.Request, param *Params) (error) {
    
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    (*param)["Token"] = token


    return nil
}

func retriveInsertedRegisterData (w http.ResponseWriter,r *http.Request,param *Params) (error) {


	return errors.New("Not correct login data")
}

func processRegisterData (w http.ResponseWriter, r *http.Request, param *Params) (error) {
    
	return nil
}