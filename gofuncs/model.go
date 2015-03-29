package hallo

import (
    "time"
    "code.google.com/p/go.crypto/bcrypt"
)

type User struct {
    Username string
    Password []byte
}

func (u *User) SetPassword( password string){
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	u.Password = hpass
}

type Date struct {
    User string
    Content string
    Date    time.Time
}

type Page struct {
    Title string
    Token string
    Content  []byte
}

type Params map[string] interface {}


