package models

type User struct {
	ID           int    `json:"id"`
	Uname        string `json:"uname"`
	Password     string `json:"pass"`
	Mail         string `json:"mail"`
	WalletAdress string `json:"wallet_adress"`
}

var DummyUsers = []User{
	{ID: 1, Uname: "sa", Password: "123", Mail: "sa@gmail.com", WalletAdress: "0x123dsadsdasd"},
	{ID: 2, Uname: "as", Password: "321", Mail: "as@gmail.com", WalletAdress: "0x321skadlkfna"},
	{ID: 3, Uname: "da", Password: "432", Mail: "da@gmail.com", WalletAdress: "0x91230sf90sd9s"},
	{ID: 4, Uname: "ad", Password: "234", Mail: "ad@gmail.com", WalletAdress: "0xs8df7sd9f87sd"},
	{ID: 5, Uname: "fa", Password: "543", Mail: "fa@gmail.com", WalletAdress: "0x687dsf6s7d8f"},
	{ID: 6, Uname: "af", Password: "345", Mail: "sa@gmail.com", WalletAdress: "0x67sd8f6ds78f6sd"},
}
