package main

import (
	"net/http"

	"batuhand.com/api/controllers"
)

func main() {

	rootName := "/api"
	defer http.ListenAndServe(":8080", nil)

	// ------------------------- USER ENDPOINTS -------------------------

	// Get all users
	// Method: GET - /get_users/
	http.HandleFunc(rootName+"/user/get_users/", controllers.GetAllUsers)

	// Get user by id
	// Method: GET - /user/%id
	http.HandleFunc(rootName+"/user/get_user/", controllers.GetUser)

	// Update user by id
	// Method: PUT - /user/%id
	http.HandleFunc(rootName+"/user/update_user/", controllers.UpdateUser)

	// Delete user by id
	// Method: DELETE - /user/%id
	http.HandleFunc(rootName+"/user/delete_user/", controllers.DeleteUser)

	// New user
	// Method: POST - /create_user/
	http.HandleFunc(rootName+"/user/create_user/", controllers.CreateUser)


	// ------------------------- TRANSACTION ENDPOINTS -------------------------

	// Get all transactions
	// Method: GET - /get_transactions/
	http.HandleFunc(rootName+"/transaction/get_transactions/", controllers.GetAllTransactions)

	// New transaction
	// Method: POST - /send_coin/
	http.HandleFunc(rootName+"/transaction/send_coin/", controllers.SendCoin)

	// Get wallet currency
	// Method: GET - /get_wallet_currency/
	http.HandleFunc(rootName+"/transaction/get_wallet_currency/", controllers.GetCurrency)

}
