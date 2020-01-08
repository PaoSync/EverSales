package controllers

import (
	"../models"
	u "../utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func sucmitRatind(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, os.Getenv("cookie_name"))
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		resp := u.Message(false, "NO user logged in")
		u.Respond(w, resp)
	} else {
		if session.Values["userRole"].(int) != 0{
			resp := u.Message(false, "User doesnt have the correct role")
			u.Respond(w, resp)
		}else {
			user := session.Values["userID"]
			rating := &models.Rating{}

			err := json.NewDecoder(r.Body).Decode(rating)
			if err != nil {
				u.Respond(w, u.Message(false, "Error while decoding request body"))
				fmt.Println(err)
				return
			}

			rating.UserID = user.(uint)
			resp := rating.Create()
			u.Respond(w, resp)
		}
	}
}