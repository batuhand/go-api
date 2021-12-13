package main

import (
	"net/http"

	"batuhand.com/api/controllers"
	"github.com/gorilla/mux"
)

func main() {

	rootName := "/api"
	router := mux.NewRouter()

	// ------------------------- USER ENDPOINTS -------------------------

	// Get all users
	// Method: GET - /get_users/
	router.HandleFunc(rootName+"/user/get_users/", controllers.GetAllUsers).Methods("GET")

	// Get user by id
	// Method: GET - /user/%id
	router.HandleFunc(rootName+"/user/get_user/", controllers.GetUser).Methods("GET")

	// Update user by id
	// Method: PUT - /user/%id
	router.HandleFunc(rootName+"/user/update_user/", controllers.UpdateUser).Methods("PUT")

	// Delete user by id
	// Method: DELETE - /user/%id
	router.HandleFunc(rootName+"/user/delete_user/", controllers.DeleteUser).Methods("DELETE")

	// New user
	// Method: POST - /create_user/
	router.HandleFunc(rootName+"/user/create_user/", controllers.CreateUser).Methods("POST")

	// ------------------------- TRANSACTION ENDPOINTS -------------------------

	// Get all transactions
	// Method: GET - /get_transactions/
	router.HandleFunc(rootName+"/transaction/get_transactions/", controllers.GetAllTransactions).Methods("GET")

	// New transaction
	// Method: POST - /send_coin/
	router.HandleFunc(rootName+"/transaction/send_coin/", controllers.SendCoin).Methods("POST")

	// Get wallet currency
	// Method: GET - /get_wallet_currency/
	router.HandleFunc(rootName+"/transaction/get_wallet_currency/", controllers.GetCurrency).Methods("GET")

	http.ListenAndServe(":8080", router)

}
