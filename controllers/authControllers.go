package controllers

import (
	"../models"
	u "../utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)
var (
	key = []byte(os.Getenv("session_password"))
	store = sessions.NewCookieStore(key)
)
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request){
	account := &models.Account{}
	//fmt.Print(r.Context())
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp:= models.Login(account.Email, account.Password)
	if resp["status"].(bool) {
		session, _ := store.Get(r, os.Getenv("cookie_name"))
		session.Values["authenticated"] = true
		session.Values["userID"] = resp["account"].(*models.Account).ID
		session.Values["userRole"] = resp["account"].(*models.Account).Role
		err = session.Save(r, w)
		if err!=nil{
			fmt.Println(err)
		}
	}
	u.Respond(w,resp)
}

func Logout(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	session.Values["authenticated"] = false
	session.Values["userID"] = ""
	session.Values["userRole"] = ""
	err := session.Save(r, w)
	if err!=nil{
		fmt.Println(err)
	}
	resp := u.Message(true, "User logged out")
	u.Respond(w,resp)
}