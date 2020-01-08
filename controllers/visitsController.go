package controllers

import (
	"../models"
	u "../utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

func NewVisit(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		resp := u.Message(false, "NO user logged in")
		u.Respond(w, resp)
	} else {
		if session.Values["userRole"].(int) != 0{
			resp := u.Message(false, "User doesnt have the correct role")
			u.Respond(w, resp)
		}else {
			vars := mux.Vars(r)
			key := vars["propertyID"]
			propertyID, err := strconv.Atoi(key)
			if err != nil {
				fmt.Println(err)
			}
			resp := models.RequestVisit(session.Values["userID"].(uint),uint(propertyID))
			u.Respond(w, resp)
		}
	}
}

func GetVisitsForUser(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		resp := u.Message(false, "NO user logged in")
		u.Respond(w, resp)
	} else {
		if session.Values["userRole"].(int) != 1{
			resp := u.Message(false, "User doesnt have the correct role")
			u.Respond(w, resp)
		}else {
			user := session.Values["userID"]
			data := models.GetVisitsForUser(user.(uint))
			resp := u.Message(true, "success")
			resp["data"] = data
			u.Respond(w, resp)
		}
	}
}

func GetVisitsForProperty(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		resp := u.Message(false, "NO user logged in")
		u.Respond(w, resp)
	} else {
		if session.Values["userRole"].(int) != 1{
			resp := u.Message(false, "User doesnt have the correct role")
			u.Respond(w, resp)
		}else {
			vars := mux.Vars(r)
			key := vars["id"]
			propertyID, err := strconv.Atoi(key)
			if err != nil {
				fmt.Println(err)
			}
			data := models.GetVisitsForProperty(uint(propertyID))
			resp := u.Message(true, "success")
			resp["data"] = data
			u.Respond(w, resp)
		}
	}
}

func ApproveVisit(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		resp := u.Message(false, "NO user logged in")
		u.Respond(w, resp)
	} else {
		if session.Values["userRole"].(int) != 1{
			resp := u.Message(false, "User doesnt have the correct role")
			u.Respond(w, resp)
		}else {
			vars := mux.Vars(r)
			key := vars["visitID"]
			visitID, err := strconv.Atoi(key)
			if err != nil {
				fmt.Println(err)
			}
			resp := models.ApproveVisit(uint(visitID))
			u.Respond(w, resp)
		}
	}
}

func DenyVisit(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		resp := u.Message(false, "NO user logged in")
		u.Respond(w, resp)
	} else {
		if session.Values["userRole"].(int) != 1{
			resp := u.Message(false, "User doesnt have the correct role")
			u.Respond(w, resp)
		}else {
			vars := mux.Vars(r)
			key := vars["visitID"]
			visitID, err := strconv.Atoi(key)
			if err != nil {
				fmt.Println(err)
			}
			resp := models.DenyVisit(uint(visitID))
			u.Respond(w, resp)
		}
	}
}