package clntservice

import (
	"cashtransfer/dev/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func VerifyUser(username string) (user utils.User, err error) {
	response, err := http.Get(utils.SERVERENDPOINT + "/user/" + username)
	if err != nil {
		return user, errors.New("User does not exist")
	}

	respBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return user, errors.New("User does not exist")
	}

	jsonErr := json.Unmarshal(respBody, &user)
	if jsonErr != nil {
		return user, errors.New("User does not exist")
	}

	if username != user.Username {
		return user, errors.New("User does not exist")
	}

	return user, err
}

func RegisterCashTransferAccount(user utils.User) (err error) {
	url := utils.SERVERENDPOINT + "/user"
	body := &user
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	return nil
}

func CreatePayPalAccount(paypalAccount utils.PayPalAccount) (err error) {
	url := utils.PAYPALENDPOINT + "/account"
	body := &paypalAccount
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	return nil
}

func CreateMoMoAccount(momoAccount utils.MomoAccount) (err error) {
	url := utils.MOMOENDPOINT + "/account"
	body := &momoAccount
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	return nil
}

func GetUserMessage(username string) (message string, err error) {
	user, err := VerifyUser(username)
	if err != nil {
		return message, err
	}
	message = user.TransferMessage.Request

	return message, err
}

func InitiateTransfer(username string, amount float32) (err error) {
	var request utils.TransferRequest
	request.Username = username
	request.Amount = amount

	url := utils.SERVERENDPOINT + "/transfer/initiate"
	body := &request
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	return nil
}

func FulfillTransfer(username string) (err error) {
	user, err := VerifyUser(username)
	if err != nil {
		return err
	}

	senderUserName := user.TransferMessage.SenderUsername
	amountString, _ := strconv.ParseFloat(strings.Fields(user.TransferMessage.Request)[1], 32)
	amount := float32(amountString)

	var request utils.TransferRequest
	request.Username = senderUserName
	request.Amount = amount

	url := utils.SERVERENDPOINT + "/transfer/fulfill/" + username
	body := &request
	err = utils.CallRestAPI("POST", url, body)
	if err != nil {
		return err
	}

	return nil
}

func GetUserBalances(username string) (momoAccount utils.MomoAccount, paypalAccount utils.PayPalAccount, err error) {

	// Get user PayPal and MoMo ID
	user, err := VerifyUser(username)
	if err != nil {
		return momoAccount, paypalAccount, err
	}

	paypalID := user.PayPalID
	momoID := user.MomoID

	// Get MoMo balance
	response, err := http.Get(utils.MOMOENDPOINT + "/account/" + momoID)
	if err != nil {
		return momoAccount, paypalAccount, errors.New("Error obtaining balances")
	}

	respBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return momoAccount, paypalAccount, errors.New("Error obtaining balances")
	}

	jsonErr := json.Unmarshal(respBody, &momoAccount)
	if jsonErr != nil {
		return momoAccount, paypalAccount, errors.New("Error obtaining balances")
	}

	// Get PayPal balance
	response, err = http.Get(utils.PAYPALENDPOINT + "/account/" + paypalID)
	if err != nil {
		return momoAccount, paypalAccount, errors.New("Error obtaining balances")
	}

	respBody, _ = ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return momoAccount, paypalAccount, errors.New("Error obtaining balances")
	}

	jsonErr = json.Unmarshal(respBody, &paypalAccount)
	if jsonErr != nil {
		return momoAccount, paypalAccount, errors.New("Error obtaining balances")
	}

	return momoAccount, paypalAccount, nil
}
