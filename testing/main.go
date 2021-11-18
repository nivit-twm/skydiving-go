package main

import (
	"bank-testing/account"
	"bank-testing/credit"
	"fmt"
	"log"
)

func main() {
	creditService := credit.NewCreditService([]string{"John"})

	accA, err := account.NewAccount(creditService, "John", "Doe", 100)
	if err != nil {
		log.Fatal(err)
	}

	_, err = accA.Deposit(100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accA.String())
}
