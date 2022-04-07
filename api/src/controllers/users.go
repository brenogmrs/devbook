package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	userID, err := repository.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("inserted ID: %d", userID)))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("get all users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("get user by id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("delete user"))
}
