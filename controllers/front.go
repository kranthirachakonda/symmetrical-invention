package controllers

import "net/http"

func RegisterController() {
	uc := newuserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
