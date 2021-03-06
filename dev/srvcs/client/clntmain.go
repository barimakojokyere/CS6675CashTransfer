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
		fmt.Println("8. Check accounts balance")
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
			fmt.Println("Please enter your full name:")
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
		case 3:
			var paypalAccount utils.PayPalAccount
			fmt.Println("Please enter your PayPal username:")
			scanner.Scan()
			paypalAccount.PayPalID = scanner.Text()
			fmt.Println("Please enter your full name:")
			scanner.Scan()
			paypalAccount.PayPalName = scanner.Text()
			fmt.Println("Please enter the amount for your PayPal account:")
			scanner.Scan()
			amountString := scanner.Text()
			amount, _ := strconv.ParseFloat(amountString, 32)
			paypalAccount.PayPalBalance = float32(amount)
			err := clntservice.CreatePayPalAccount(paypalAccount)
			if err != nil {
				fmt.Println("Could not create PayPal account.")
				continue
			}
			fmt.Println("PayPal account successfully created!!!")
			continue
		case 4:
			var momoAccount utils.MomoAccount
			fmt.Println("Please enter your MTN Mobile Money username:")
			scanner.Scan()
			momoAccount.MomoID = scanner.Text()
			fmt.Println("Please enter your full name:")
			scanner.Scan()
			momoAccount.MomoName = scanner.Text()
			fmt.Println("Please enter the amount for your MTN Mobile Money account:")
			scanner.Scan()
			amountString := scanner.Text()
			amount, _ := strconv.ParseFloat(amountString, 32)
			momoAccount.MomoBalance = float32(amount)
			err := clntservice.CreateMoMoAccount(momoAccount)
			if err != nil {
				fmt.Println("Could not create MTN Mobile Money account.")
				continue
			}
			fmt.Println("MoMo account successfully created!!!")
			continue
		case 5:
			if username != "" {
				message, err := clntservice.GetUserMessage(username)
				if err != nil {
					fmt.Println("Could not retrieve your message")
					continue
				}
				if message != "" {
					fmt.Println("Message: " + message)
				} else {
					fmt.Println("You have no messages at this time.")
				}
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
			} else {
				fmt.Println("Please log in first...")
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
			} else {
				fmt.Println("Please log in first...")
			}
			continue
		case 8:
			if username != "" {
				momoAccnt, paypalAccnt, err := clntservice.GetUserBalances(username)
				if err != nil {
					fmt.Println("Could not obtain user balances")
					continue
				}
				fmt.Println("PayPal Balance: " + fmt.Sprint(paypalAccnt.PayPalBalance))
				fmt.Println("MTN Mobile Money Balance: " + fmt.Sprint(momoAccnt.MomoBalance))
			} else {
				fmt.Println("Please log in first...")
			}
			continue
		default:
			fmt.Println("You entered an incorrect value. Please enter a correct one.")
			continue
		}
	}
}
