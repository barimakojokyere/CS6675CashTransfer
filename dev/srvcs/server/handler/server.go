package handler

import (
	"cashtransfer/dev/srvcs/server/svrsrvc"
	"cashtransfer/dev/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user utils.User
	json.Unmarshal(reqBody, &user)
	err := svrsrvc.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := svrsrvc.RetrieveUser(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	err := svrsrvc.DeleteUser(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(username)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var user utils.User
	json.Unmarshal(reqBody, &user)
	svrsrvc.UpdateUser(username, user)
	json.NewEncoder(w).Encode(user)
}
