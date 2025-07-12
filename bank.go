package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func run_bank() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("============ERROR===============")
		fmt.Println(err.Error())
		fmt.Println("================================")
		panic("")
	}

	for {
		fmt.Println("Welcome To Bank")
		fmt.Println("What do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			{
				fmt.Print("Your deposit : ")
				var depositAmount float64
				fmt.Scan(&depositAmount)
				if depositAmount <= 0 {
					fmt.Println("Must Be Greater Than 0")
					continue
				}
				accountBalance = accountBalance + depositAmount
				fmt.Println("Your balance updated : ", accountBalance)
				writeBalanceToBank(accountBalance)
			}
		case 3:
			{
				fmt.Print("Withdrawal amount : ")
				var withdrawal float64
				fmt.Scan(&withdrawal)

				if withdrawal <= 0 {
					fmt.Println("Must Be Greater Than 0")
					continue
				}

				if withdrawal > accountBalance {
					fmt.Println("Invalid Amount")
					continue
				}
				accountBalance = accountBalance - withdrawal
				fmt.Println("Your balance updated : ", accountBalance)
				writeBalanceToBank(accountBalance)
			}
		default:
			{
				fmt.Println("Good bye")
				break
			}
		}
	}

}

func writeBalanceToBank(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 1000, errors.New("Failed To Find File")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed To Convert To Float Value")
	}
	return balance, nil
}
