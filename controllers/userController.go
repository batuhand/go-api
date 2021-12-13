package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"batuhand.com/api/models"
)

// Returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(models.DummyUsers)
	CheckError(err)
	fmt.Println(string(users))
	fmt.Fprint(w, string(users))

}

// Returns spesific user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url_list[4])
	CheckError(err)
	for i, user := range models.DummyUsers {
		fmt.Println(i, " ", user.ID)
		if id == user.ID {
			user, err := json.Marshal(user)
			CheckError(err)
			fmt.Fprint(w, string(user))
		}
	}

}

// Updates user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url_list[4])
	CheckError(err)
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	newUser := models.User{}
	json.Unmarshal(body, &newUser)
	for i, user := range models.DummyUsers {
		fmt.Println(i, " ", user.ID)
		if id == user.ID {
			if newUser.Uname != "" {
				models.DummyUsers[i].Uname = newUser.Uname
			}
			if newUser.Password != "" {
				models.DummyUsers[i].Password = newUser.Password
			}
			if newUser.Mail != "" {
				models.DummyUsers[i].Mail = newUser.Mail
			}
			if newUser.WalletAdress != "" {
				models.DummyUsers[i].WalletAdress = newUser.WalletAdress
			}
		}
	}

}

// For removing object from the given index
func RemoveIndex(users []models.User, index int) []models.User {
	return append(users[:index], users[index+1:]...)
}

// Deletes user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url_list[4])
	CheckError(err)
	for i, user := range models.DummyUsers {
		fmt.Println(i, " ", user.ID)
		if id == user.ID {
			models.DummyUsers = RemoveIndex(models.DummyUsers, i)
			fmt.Fprint(w, "User deleted")
		}
	}

}

// Creates user and appends it to the dummy data array
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	user := models.User{}
	json.Unmarshal(body, &user)
	models.DummyUsers = append(models.DummyUsers, user)
	//CheckError(err)
	fmt.Fprint(w, "User created")

}
