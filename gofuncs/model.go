package hallo

import (
    "time"
)

type Params map[string] interface {}

type Date struct {
    Person  string
    Content string
    Date    time.Time
}

type Person struct {
    Name string
    Password string
}

type Page struct {
    Title string
    Token string
    Content  []byte
}

