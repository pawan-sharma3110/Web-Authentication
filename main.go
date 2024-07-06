package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var P1 = person{
	FirstName: "Pawan",
	LastName:  "Sharma",
}
var P3 = person{
	FirstName: "Ashu",
	LastName:  "Kala",
}
var P2 = person{
	FirstName: "Ravi",
	LastName:  "Sharma",
}
var all = []person{
	P1, P2, P3,
}

func main() {
	// http.HandleFunc("/new/person", NewPerson)
	// http.HandleFunc("/all/person", allPerson)
	// http.ListenAndServe(":8080", nil)
	pass := "12345678"
	pass2 := "798832"
	hashedPas, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}
	err = compareHashPass(hashedPas, []byte(pass2))
	if err != nil {
		log.Fatalf("Not loged in")
	}
	log.Println("Loged In")
}

func hashPassword(password string) ([]byte, error) {
	hasedpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generateing bcrypt hash from password %w", err)
	}
	return hasedpass, nil
}

func compareHashPass(hashPass, password []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password %w", err)
	}
	return nil
}

// func NewPerson(w http.ResponseWriter, r *http.Request) {
// 	var newPerson person
// 	json.NewDecoder(r.Body).Decode(&newPerson)
// 	all = append(all, newPerson)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	err := json.NewEncoder(w).Encode(all)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func allPerson(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	err := json.NewEncoder(w).Encode(all)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
