package main

import (
	"fmt"

	"example.com/course/fileops"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"

func run_bank() {
	var accountBalance, err = fileops.GetFloatFromFile(ACCOUNT_BALANCE_FILE)

	if err != nil {
		fmt.Println("============ERROR===============")
		fmt.Println(err.Error())
		fmt.Println("================================")
		panic("")
	}

	for {
		presentOptions()

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
				fileops.WriteFloatToFile(accountBalance, ACCOUNT_BALANCE_FILE)
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
				fileops.WriteFloatToFile(accountBalance, ACCOUNT_BALANCE_FILE)
			}
		default:
			{
				fmt.Println("Good bye")
				return
			}
		}
	}

}
