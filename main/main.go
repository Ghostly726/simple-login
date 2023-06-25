package main

import (
	"net/http"
	"timeserver"
)

func main() {
	http.HandleFunc("/", timeserver.ServeLoginHTML)
	http.HandleFunc("/authenticate", timeserver.AuthenticateUser)
	http.HandleFunc("/signup", timeserver.ServeSignupHTML)
	http.HandleFunc("/appendUser", timeserver.AppendUser)

	http.ListenAndServe(":8080", nil)
}
