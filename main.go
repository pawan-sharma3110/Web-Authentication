package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	p1 := person{
		FirstName: "Pawan",
		LastName:  "Sharma",
	}
	p3 := person{
		FirstName: "Ashu",
		LastName:  "Kala",
	}
	p2 := person{
		FirstName: "Ravi",
		LastName:  "Sharma",
	}
	all := []person{
		p1, p2, p3,
	}
	res, err := json.Marshal(all)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshaled data:", string(res))

	detail := []person{}
	err = json.Unmarshal(res, &detail)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unmarshaled data:", detail)
}
