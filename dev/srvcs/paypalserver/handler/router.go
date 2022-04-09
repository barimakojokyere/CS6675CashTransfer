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
	router.HandleFunc(utils.PAYPALRESTAPIBASUEURL+"/account", CreateAccount).Methods("POST")
	router.HandleFunc(utils.PAYPALRESTAPIBASUEURL+"/account/{username}", RetrieveAccount).Methods("GET")
	router.HandleFunc(utils.PAYPALRESTAPIBASUEURL+"/transfer/{username}", MakeTransfer).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
