package svrsrvc

import (
	"cashtransfer/dev/srvcs/database"
	"cashtransfer/dev/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user utils.User) (err error) {
	err = database.InsertIntoDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, user)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveUser(username string) (user utils.User, err error) {
	output, err := database.RetrieveFromDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username)
	if err != nil {
		return user, err
	}

	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &user)
	return user, nil
}

func DeleteUser(username string) (err error) {
	err = database.RemoveFromDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(username string, user utils.User) (err error) {
	err = database.UpdateInDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username, user)
	if err != nil {
		return err
	}

	return nil
}

func InitiateTransfer(transfer utils.TransferRequest) (err error) {

	// Create transfer message
	transfer.Completed = false

	err = database.InsertIntoDB(utils.CASHTRANSFERDBNAME, utils.TRANSFERREQCOLLECTION, transfer)
	if err != nil {
		return err
	}

	var message utils.TransferMessage
	message.SenderUsername = transfer.Username
	message.Request = "Send " + fmt.Sprint(transfer.Amount)

	var users []utils.User

	output, err := database.RetrieveAllInCollection(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION)
	if err != nil {
		return err
	}

	var user utils.User

	for _, currOutput := range output {
		outputBytes, _ := bson.Marshal(currOutput)
		bson.Unmarshal(outputBytes, &user)
		users = append(users, user)
	}

	for _, currUser := range users {
		if currUser.TransferMessage == (utils.TransferMessage{}) && currUser.Username != transfer.Username {
			currUser.TransferMessage = message
			UpdateUser(currUser.Username, currUser)
		}
	}

	return nil
}

func MakeMomoTransfer(senderUsername, receiverUserName string, amount float32) (err error) {

	var momoTransferInfo utils.TransferInfo
	momoTransferInfo.DestUserID = receiverUserName
	momoTransferInfo.TransferAmount = amount

	url := utils.MOMOENDPOINT + "/transfer/" + senderUsername
	body := &momoTransferInfo
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	return nil
}

func MakePayPalTransfer(senderUsername, receiverUserName string, amount float32) (err error) {
	return nil
}

func FulfillTransfer(crowdUserUsername string, transfer utils.TransferRequest) (err error) {

	var senderPayPalID string
	var crowdUserPayPalID string
	var senderMoMoID string
	var crowdUserMoMoID string

	// Find momo and paypal usernames
	var users []utils.User

	output, err := database.RetrieveAllInCollection(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION)
	if err != nil {
		return err
	}

	var user utils.User

	for _, currOutput := range output {
		outputBytes, _ := bson.Marshal(currOutput)
		bson.Unmarshal(outputBytes, &user)
		users = append(users, user)
	}

	for _, currUser := range users {
		if currUser.Username == crowdUserUsername {
			crowdUserMoMoID = currUser.MomoID
			crowdUserPayPalID = currUser.PayPalID
		}
		if currUser.Username == transfer.Username {
			senderMoMoID = currUser.MomoID
			senderPayPalID = currUser.PayPalID
		}
	}

	//Do MoMo transfer
	err = MakeMomoTransfer(senderMoMoID, crowdUserMoMoID, transfer.Amount)
	if err != nil {
		return err
	}

	//Do PayPal transfer
	err = MakePayPalTransfer(senderPayPalID, crowdUserPayPalID, transfer.Amount)
	if err != nil {
		return err
	}

	return nil
}
