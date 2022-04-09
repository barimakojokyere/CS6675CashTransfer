package momosrvc

import (
	"cashtransfer/dev/srvcs/database"
	"cashtransfer/dev/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount(account utils.MomoAccount) (err error) {
	err = database.InsertIntoDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, account)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveAccount(username string) (account utils.MomoAccount, err error) {
	output, err := database.RetrieveFromDB(utils.MOMODBNAME, utils.ACCNTSCOLLECTION, username)
	if err != nil {
		return account, err
	}

	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &account)
	return account, nil
}

func MakeTransfer(username string, transferInfo utils.TransferInfo) (err error) {

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

	return nil
}
