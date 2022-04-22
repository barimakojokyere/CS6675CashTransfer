package paypalsrvc

import (
	"cashtransfer/dev/srvcs/database"
	"cashtransfer/dev/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount(account utils.PayPalAccount) (err error) {
	fmt.Println("Creating PayPal account.")
	err = database.InsertIntoDB(utils.PAYPALDBNAME, utils.ACCNTSCOLLECTION, account)
	if err != nil {
		return err
	}
	fmt.Println("PayPal account successfully created.")
	return nil
}

func RetrieveAccount(username string) (account utils.PayPalAccount, err error) {

	fmt.Println("Retrieving PayPal account...")

	output, err := database.RetrieveFromDB(utils.PAYPALDBNAME, utils.ACCNTSCOLLECTION, username)
	if err != nil {
		return account, err
	}

	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &account)

	fmt.Println("PayPal account successfully retrieved.")

	return account, nil
}

func RetrieveAllAccounts() (accounts []utils.PayPalAccount, err error) {

	fmt.Println("Retrieving all PayPal accounts")

	output, err := database.RetrieveAllInCollection(utils.PAYPALDBNAME, utils.ACCNTSCOLLECTION)
	if err != nil {
		return accounts, err
	}

	var currAccount utils.PayPalAccount

	for _, currOutput := range output {
		outputBytes, _ := bson.Marshal(currOutput)
		bson.Unmarshal(outputBytes, &currAccount)
		accounts = append(accounts, currAccount)
	}

	fmt.Println("All PayPal accounts have been successfully retrieved.")

	return accounts, nil
}

func MakeTransfer(username string, transferInfo utils.TransferInfo) (err error) {

	fmt.Println("Making PayPal transfer...")

	var destUserAccount utils.PayPalAccount

	// increase destination users amount
	output, err := database.RetrieveFromDB(utils.PAYPALDBNAME, utils.ACCNTSCOLLECTION, transferInfo.DestUserID)
	if err != nil {
		return err
	}
	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &destUserAccount)

	destUserAccount.PayPalBalance = destUserAccount.PayPalBalance + transferInfo.TransferAmount

	err = database.UpdateInDB(utils.PAYPALDBNAME, utils.ACCNTSCOLLECTION, transferInfo.DestUserID, destUserAccount)
	if err != nil {
		return err
	}

	var sendingUserAccount utils.PayPalAccount

	// decrease sending users amount
	output, err = database.RetrieveFromDB(utils.PAYPALDBNAME, utils.ACCNTSCOLLECTION, username)
	if err != nil {
		return err
	}

	outputBytes, _ = bson.Marshal(output)
	bson.Unmarshal(outputBytes, &sendingUserAccount)

	sendingUserAccount.PayPalBalance = sendingUserAccount.PayPalBalance - transferInfo.TransferAmount

	err = database.UpdateInDB(utils.PAYPALDBNAME, utils.ACCNTSCOLLECTION, username, sendingUserAccount)
	if err != nil {
		return err
	}

	fmt.Println("PayPal transfer was successful...")

	return nil
}
