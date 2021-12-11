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
	if r.Method == "GET" {
		users, err := json.Marshal(models.DummyUsers)
		CheckError(err)
		fmt.Println(string(users))
		fmt.Fprint(w, string(users))
	} else {
		fmt.Fprintf(w, "Bad Request")
	}

}

// Returns spesific user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
	}

}


// Updates user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
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
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
	}

}

// For removing object from the given index
func RemoveIndex(users []models.User, index int) []models.User {
	return append(users[:index], users[index+1:]...)
}


// Deletes user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
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

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
	}

}

// Creates user and appends it to the dummy data array
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		CheckError(err)
		user := models.User{}
		json.Unmarshal(body, &user)
		models.DummyUsers = append(models.DummyUsers, user)
		//CheckError(err)
		fmt.Fprint(w, "User created")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
	}

}
