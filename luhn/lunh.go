package main

import (
	"fmt"

	creditcard "github.com/durango/go-credit-card"
)

func main() {
	card := creditcard.Card{Number: "4426349636241779", Cvv: "588", Month: "02", Year: "2019"}
	err := card.Validate()
	fmt.Println(err)
	fmt.Println(card.Company.Long)
	fmt.Println(card.Company.Short)
}
