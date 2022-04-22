package momosrvc

import (
	"cashtransfer/dev/srvcs/database"
	"cashtransfer/dev/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount(account utils.MomoAccount) (err error) {
	fmt.Println("Creating MoMo account...")
	err = database.InsertIntoDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, account)
	if err != nil {
		return err
	}
	fmt.Println("MoMo account successfully created.")
	return nil
}

func RetrieveAccount(username string) (account utils.MomoAccount, err error) {

	fmt.Println("Retrieving MoMo account...")

	output, err := database.RetrieveFromDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, username)
	if err != nil {
		return account, err
	}

	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &account)

	fmt.Println("MoMo account successfully retrieved.")

	return account, nil
}

func RetrieveAllAccounts() (accounts []utils.MomoAccount, err error) {

	fmt.Println("Retrieving all MoMo accounts")

	output, err := database.RetrieveAllInCollection(utils.MOMODBNAME, utils.ACCNTSCOLLECTION)
	if err != nil {
		return accounts, err
	}

	var currAccount utils.MomoAccount

	for _, currOutput := range output {
		outputBytes, _ := bson.Marshal(currOutput)
		bson.Unmarshal(outputBytes, &currAccount)
		accounts = append(accounts, currAccount)
	}

	fmt.Println("All MoMo accounts have been successfully retrieved.")

	return accounts, nil
}

func MakeTransfer(username string, transferInfo utils.TransferInfo) (err error) {

	fmt.Println("Making MoMo transfer...")

	var destUserAccount utils.MomoAccount

	// increase destination users amount
	output, err := database.RetrieveFromDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, transferInfo.DestUserID)
	if err != nil {
		return err
	}
	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &destUserAccount)

	destUserAccount.MomoBalance = destUserAccount.MomoBalance + transferInfo.TransferAmount

	err = database.UpdateInDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, transferInfo.DestUserID, destUserAccount)
	if err != nil {
		return err
	}

	var sendingUserAccount utils.MomoAccount

	// decrease sending users amount
	output, err = database.RetrieveFromDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, username)
	if err != nil {
		return err
	}

	outputBytes, _ = bson.Marshal(output)
	bson.Unmarshal(outputBytes, &sendingUserAccount)

	sendingUserAccount.MomoBalance = sendingUserAccount.MomoBalance - transferInfo.TransferAmount

	err = database.UpdateInDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, username, sendingUserAccount)
	if err != nil {
		return err
	}

	fmt.Println("MoMo transfer was successful...")

	return nil
}
