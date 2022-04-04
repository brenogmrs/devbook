package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Create user"))
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
