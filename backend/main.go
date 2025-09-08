package main

import (
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

const (
	SessionKey  = "session-key"
	SessionName = "session-name"
)


func main() {

}