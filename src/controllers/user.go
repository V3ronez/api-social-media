package controllers

import "net/http"

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all user!"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}
