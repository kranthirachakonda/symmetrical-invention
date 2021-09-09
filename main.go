package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
	fmt.Println("Hello Gopher!!")
	u := models.User{
		ID:        2,
		FirstName: "Kraz",
		LastName:  "Rach",
	}
	fmt.Println(u)
}
