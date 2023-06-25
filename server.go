package timeserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ServeLoginHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "/home/ghost/ser1/templates/index.html")
		fmt.Println(http.StatusOK)
	} else {
		fmt.Println(http.StatusBadRequest)
	}
}

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

//* Begginer user database, sign up form coming in
var ValidUsers = []User{
	{Email: "HIO@gmail.com", Name: "HIO", Password: "ab123"},
	{Email: "AB2@gmail.com", Name: "AB2", Password: "2007ytw"},
	{Email: "IO2@gmail.com", Name: "IO2", Password: "po568/fuey@124"},
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var trialedUser User

	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&trialedUser)

		if err != nil {
			fmt.Println("Error in parsing JSON to golang")
		}

		if CorrelateUser(&trialedUser, ValidUsers) {
			fmt.Println(trialedUser.Name + " " + trialedUser.Email + " " + trialedUser.Password)
		} else {
			log.Fatal("User Invalid")
		}

	} else {
		fmt.Println(http.StatusBadRequest)
	}
}

func ServeSignupHTML(w http.ResponseWriter, r *http.Request) {

}

func AppendUser(w http.ResponseWriter, r *http.Request) {

}
