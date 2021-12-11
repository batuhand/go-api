package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"batuhand.com/api/models"
)

// Returns all the transactions
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		transactions, err := json.Marshal(models.DummyTransactions)
		CheckError(err)
		fmt.Println(string(transactions))
		fmt.Fprint(w, string(transactions))
	} else {
		fmt.Fprintf(w, "Bad Request")
	}
}

// Creates a transaction if there are enough currency in the sender wallet adress
func SendCoin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		CheckError(err)
		transaction := models.Transaction{}
		json.Unmarshal(body, &transaction)
		senderCurrency := GetWalletAmount(transaction.SenderWalletAdress)
		if senderCurrency > transaction.Amount {
			models.DummyTransactions = append(models.DummyTransactions, transaction)
			//CheckError(err)
			fmt.Fprint(w, "Transaction created")
		} else {
			fmt.Fprintf(w, "Transaction failed, not enough currency")
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
	}

}

// Returns the currency of the given wallet adress
func GetCurrency(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url_list := strings.Split(r.URL.Path, "/")
		wallet_adress := url_list[4]
		currency := GetWalletAmount(wallet_adress)
		fmt.Fprintf(w, "%f", currency)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
	}

}

// Returns the currency of the given wallet
func GetWalletAmount(walletId string) float64 {
	var currency float64
	for _, t := range models.DummyTransactions {
		if t.RecieverWalletAdress == walletId {
			currency = currency + t.Amount
		}
		if t.SenderWalletAdress == walletId {
			currency = currency - t.Amount
		}
	}
	return currency
}
