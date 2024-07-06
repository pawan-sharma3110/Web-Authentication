package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	// res, err := json.Marshal(all)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Marshaled data:", string(res))

	// detail := []person{}
	// err = json.Unmarshal(res, &detail)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Unmarshaled data:", detail)
	http.HandleFunc("/new/person", NewPerson)
	http.HandleFunc("/all/person", allPerson)
	http.ListenAndServe(":8080", nil)

}
func NewPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson person
	json.NewDecoder(r.Body).Decode(&newPerson)
	all = append(all, newPerson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(all)
	if err != nil {
		log.Fatal(err)
	}
}

func allPerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(all)
	if err != nil {
		log.Fatal(err)
	}
}
