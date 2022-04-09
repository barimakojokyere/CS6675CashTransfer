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
	router.HandleFunc(utils.RESTAPIBASEURL+"/user", CreateUser).Methods("POST")
	router.HandleFunc(utils.RESTAPIBASEURL+"/user/{username}", DeleteUser).Methods("DELETE")
	router.HandleFunc(utils.RESTAPIBASEURL+"/user/{username}", UpdateUser).Methods("PATCH")
	router.HandleFunc(utils.RESTAPIBASEURL+"/user/{username}", RetrieveUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
