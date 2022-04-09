package handler

import (
	"cashtransfer/dev/srvcs/paypalserver/paypalsrvc"
	"cashtransfer/dev/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account utils.PayPalAccount
	json.Unmarshal(reqBody, &account)
	fmt.Println(account)
	err := paypalsrvc.CreateAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(account)
	}
}

func RetrieveAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	account, err := paypalsrvc.RetrieveAccount(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(account)
	}
}

func RetrieveAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := paypalsrvc.RetrieveAllAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(accounts)
	}
}

func MakeTransfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var transferInfo utils.TransferInfo
	json.Unmarshal(reqBody, &transferInfo)

	err := paypalsrvc.MakeTransfer(username, transferInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(transferInfo)
	}
}
