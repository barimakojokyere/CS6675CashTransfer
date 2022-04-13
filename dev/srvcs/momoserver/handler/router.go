package handler

import (
	"cashtransfer/dev/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	// Create a router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(utils.MOMORESTAPIBASUEURL+"/account", CreateAccount).Methods("POST")
	router.HandleFunc(utils.MOMORESTAPIBASUEURL+"/account/{username}", RetrieveAccount).Methods("GET")
	router.HandleFunc(utils.MOMORESTAPIBASUEURL+"/transfer/{username}", MakeTransfer).Methods("POST")
	router.HandleFunc(utils.MOMORESTAPIBASUEURL+"/accounts", RetrieveAllAccounts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}
