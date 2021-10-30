package main

import (
	"fmt"
	"packages/auxiliary"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Learning Golang")
	auxiliary.Write()
	err := checkmail.ValidateFormat("johndoe@email.com")
	fmt.Println(err)
}
