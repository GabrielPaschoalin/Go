package main

import (
	"fmt"

	"exemple.com/bank/utils"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, error := utils.GetFloatFromFile(accountBalanceFile, 1000)

	if error != nil {
		// fmt.Println("ERROR")
		// fmt.Println(error)
		// return
		panic(error)
	}

	var depositMoney float64
	var withdrawMoney float64
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	// for i := 0; i >= 0; i++ {
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("How much do you want to deposit? ")
			fmt.Scan(&depositMoney)

			if depositMoney <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositMoney
			fmt.Println("Balance updated! New amount: ", accountBalance)
			utils.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("How much do you want to withdraw? ")
			fmt.Scan(&withdrawMoney)

			if withdrawMoney <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawMoney > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawMoney
			fmt.Println("Balance updated! New amount: ", accountBalance)
			utils.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			// break -> doesn't work
			return
		}
	}

}
