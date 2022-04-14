package main

import (
	"bufio"
	"cashtransfer/dev/srvcs/client/clntservice"
	"cashtransfer/dev/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var username string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n\n**************************************************************")
		fmt.Println("1. Login")
		fmt.Println("2. Create Account")
		fmt.Println("3. Create PayPal Account")
		fmt.Println("4. Create Mobile Money Account")
		fmt.Println("5. Check Message")
		fmt.Println("6. Make Transfer")
		fmt.Println("7. Accept Transfer Request")
		fmt.Println("**************************************************************")
		fmt.Println("Please enter the number for the action you want and hit Enter:")

		var choice int
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			fmt.Println("Please enter your username:")
			scanner.Scan()
			username = scanner.Text()
			user, err := clntservice.VerifyUser(username)
			if err != nil {
				fmt.Println("You do not have an account. Please use option 2 to create an account.")
				continue
			}
			fmt.Println("Account verified successfully. Account information:")
			fmt.Println("Username: " + user.Username)
			fmt.Println("Name: " + user.Name)
			fmt.Println("Email: " + user.Email)
			fmt.Println("Country: " + user.Country)
			continue
		case 2:
			var user utils.User
			fmt.Println("Please enter a username:")
			scanner.Scan()
			user.Username = scanner.Text()
			username = user.Username
			fmt.Println("Please enter you full name:")
			scanner.Scan()
			user.Name = scanner.Text()
			fmt.Println("Please enter your email:")
			scanner.Scan()
			user.Email = scanner.Text()
			fmt.Println("Please enter your country:")
			scanner.Scan()
			user.Country = scanner.Text()
			fmt.Println("Please enter your PayPal id:")
			scanner.Scan()
			user.PayPalID = scanner.Text()
			fmt.Println("Please enter your Mobile Money id:")
			scanner.Scan()
			user.MomoID = scanner.Text()
			err := clntservice.RegisterCashTransferAccount(user)
			if err != nil {
				fmt.Println("Could not create the account due to: " + err.Error())
				continue
			}
			fmt.Println("Account successfully created!!!")
			continue
		case 5:
			if username != "" {
				message, err := clntservice.GetUserMessage(username)
				if err != nil {
					fmt.Println("Could not retrieve your message")
					continue
				}
				fmt.Println("Message: " + message)
			} else {
				fmt.Println("Please log in first...")
			}
			continue
		case 6:
			if username != "" {
				fmt.Println("Please enter the amount you want to transfer:")
				scanner.Scan()
				amountString := scanner.Text()
				amount, _ := strconv.ParseFloat(amountString, 32)
				err := clntservice.InitiateTransfer(username, float32(amount))
				if err != nil {
					fmt.Println("Could not initiate money transfer at this time.")
					continue
				}
				fmt.Println("Money transfer has been successfully initiated.")
			}
			continue
		case 7:
			if username != "" {
				err := clntservice.FulfillTransfer(username)
				if err != nil {
					fmt.Println("Could not fulfill money transfer.")
					continue
				}
				fmt.Println("Money transfer successfully fulfilled")
			}
			continue
		}
	}
}
