package hallo


import (
    "log"
	)

func reportValue(s1 string, val interface {}) {
    log.Println(s1," ",val)
}