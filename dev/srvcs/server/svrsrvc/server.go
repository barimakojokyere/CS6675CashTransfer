package svrsrvc

import (
	"cashtransfer/dev/srvcs/database"
	"cashtransfer/dev/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user utils.User) (err error) {
	fmt.Println("Creating user account.")
	err = database.InsertIntoDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, user)
	if err != nil {
		return err
	}
	fmt.Println("User account successfully created.")
	return nil
}

func RetrieveUser(username string) (user utils.User, err error) {

	fmt.Println("Retrieving user account.")

	output, err := database.RetrieveFromDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username)
	if err != nil {
		return user, err
	}

	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &user)

	fmt.Println("User account retrieved successfully.")

	return user, nil
}

func DeleteUser(username string) (err error) {
	fmt.Println("Deleting user account.")
	err = database.RemoveFromDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username)
	if err != nil {
		return err
	}
	fmt.Println("User account successfully deleted.")
	return nil
}

func UpdateUser(username string, user utils.User) (err error) {
	fmt.Println("Updating user account.")
	err = database.UpdateInDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username, user)
	if err != nil {
		return err
	}
	fmt.Println("User account successfully updated.")
	return nil
}

func InitiateTransfer(transfer utils.TransferRequest) (err error) {

	fmt.Println("Initiating money transfer.")

	// Create transfer message
	transfer.Completed = false

	err = database.InsertIntoDB(utils.CASHTRANSFERDBNAME, utils.TRANSFERREQCOLLECTION, transfer)
	if err != nil {
		return err
	}

	var message utils.TransferMessage
	message.SenderUsername = transfer.Username
	message.Request = "Send " + fmt.Sprint(transfer.Amount)

	output, err := database.RetrieveAllInCollection(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION)
	if err != nil {
		return err
	}

	var user utils.User

	for _, currOutput := range output {
		outputBytes, _ := bson.Marshal(currOutput)
		bson.Unmarshal(outputBytes, &user)
		if user.Username != transfer.Username {
			user.TransferMessage = message
			UpdateUser(user.Username, user)
		}
	}

	fmt.Println("Money transfer successfully initiated.")

	return nil
}

func RemoveMessages() (err error) {

	output, err := database.RetrieveAllInCollection(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION)
	if err != nil {
		return err
	}

	var user utils.User

	for _, currOutput := range output {
		outputBytes, _ := bson.Marshal(currOutput)
		bson.Unmarshal(outputBytes, &user)
		user.TransferMessage = utils.TransferMessage{}
		UpdateUser(user.Username, user)
	}

	return nil
}

func MakeMomoTransfer(senderUsername, receiverUserName string, amount float32) (err error) {

	fmt.Println("Processing MoMo money transfer.")

	var momoTransferInfo utils.TransferInfo
	momoTransferInfo.DestUserID = receiverUserName
	momoTransferInfo.TransferAmount = amount

	url := utils.MOMOENDPOINT + "/transfer/" + senderUsername
	body := &momoTransferInfo
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	fmt.Println("MoMo money transfer successfully processed.")

	return nil
}

func MakePayPalTransfer(senderUsername, receiverUserName string, amount float32) (err error) {

	fmt.Println("Processing PayPal money transfer.")

	var paypalTransferInfo utils.TransferInfo
	paypalTransferInfo.DestUserID = receiverUserName
	paypalTransferInfo.TransferAmount = amount

	url := utils.PAYPALENDPOINT + "/transfer/" + senderUsername
	body := &paypalTransferInfo
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	fmt.Println("PayPal money transfer successfully processed.")

	return nil
}

func FulfillTransfer(crowdUserUsername string, transfer utils.TransferRequest) (err error) {

	fmt.Println("Processing money transfer.")

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
	err = MakePayPalTransfer(crowdUserPayPalID, senderPayPalID, transfer.Amount-(0.04*transfer.Amount))
	if err != nil {
		return err
	}

	// Remove transfer entry
	err = database.RemoveFromDB(utils.CASHTRANSFERDBNAME, utils.TRANSFERREQCOLLECTION, transfer.Username)
	if err != nil {
		return err
	}

	// Remove messages
	err = RemoveMessages()
	if err != nil {
		return err
	}

	fmt.Println("Money transfer successfully processed.")

	return nil
}
