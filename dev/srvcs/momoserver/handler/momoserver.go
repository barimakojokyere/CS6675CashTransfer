package handler

import (
	"cashtransfer/dev/srvcs/momoserver/momosrvc"
	"cashtransfer/dev/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account utils.MomoAccount
	json.Unmarshal(reqBody, &account)
	err := momosrvc.CreateAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(account)
	}
}

func RetrieveAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	account, err := momosrvc.RetrieveAccount(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(account)
	}
}

func MakeTransfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var transferInfo utils.TransferInfo
	json.Unmarshal(reqBody, &transferInfo)

	err := momosrvc.MakeTransfer(username, transferInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(transferInfo)
	}
}

func RetrieveAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := momosrvc.RetrieveAllAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(accounts)
	}
}
