package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	u "../utils"
	"os"
	"../models"
	"strconv"
)

var CreateProperty = func(w http.ResponseWriter, r *http.Request){
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
			property := &models.Property{}

			err := json.NewDecoder(r.Body).Decode(property)
			if err != nil {
				u.Respond(w, u.Message(false, "Error while decoding request body"))
				fmt.Println(err)
				return
			}

			property.UserId = user.(uint)
			resp := property.Create()
			u.Respond(w, resp)
		}
	}
}

func ActiveProperties (w http.ResponseWriter, r *http.Request){
	data := models.GetActiveProperties()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func DeletePropertyByID (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	vars := mux.Vars(r)
	key := vars["id"]
	propertyID, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
	}
	resp := models.DeleteProperty(uint(propertyID),session.Values["userID"].(uint))
	u.Respond(w,resp)
}

func TogglePropertyStatus (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	vars := mux.Vars(r)
	key := vars["id"]
	propertyID, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
	}
	resp := models.TogglePropertyStatus(uint(propertyID),session.Values["userID"].(uint))
	u.Respond(w,resp)
}

func PropertyInformation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	propertyID, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
	}
	resp := models.GetPropertyByID(uint(propertyID))
	u.Respond(w,resp)
}

func ModifyProperty(w http.ResponseWriter, r *http.Request){
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
			property := &models.Property{}

			err := json.NewDecoder(r.Body).Decode(property)
			if err != nil {
				u.Respond(w, u.Message(false, "Error while decoding request body"))
				fmt.Println(err)
				return
			}
			vars := mux.Vars(r)
			key := vars["id"]
			propertyID, err := strconv.Atoi(key)
			if err != nil {
				fmt.Println(err)
			}
			property.ID = uint(propertyID)
			resp := property.Modify(user.(uint))
			u.Respond(w, resp)
		}
	}
}

func SearchProperties(w http.ResponseWriter, r *http.Request) {
	minSize := r.FormValue("minSize")
	maxSize := r.FormValue("maxSize")
	rooms := r.FormValue("rooms")
	bathrooms := r.FormValue("bathrooms")
	parking := r.FormValue("parking")
	floors := r.FormValue("floors")
	pets := r.FormValue("pets")
	propertyType := r.FormValue("property_type")
	maxPrice := r.FormValue("maxPrice")
	minPrice := r.FormValue("minPrice")
	privateSecurity := r.FormValue("private_security")
	capacity := r.FormValue("capacity")
	data := models.Search(minSize,maxSize,rooms,bathrooms,parking,floors,pets,propertyType,maxPrice,minPrice,privateSecurity,capacity)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
