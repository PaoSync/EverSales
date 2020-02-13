package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"./models"
)
func clearTable() {
	models.GetDB().Exec("DELETE FROM accounts")
}
func TestCreateAccount(t *testing.T){
	clearTable()
	payload := []byte(`{"email" : "test5@gmail.com", "password" : "secret", "username" : "user1", "role" : 0}`)
	r,_:= http.NewRequest("POST", "/api/user/new", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w,r)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.Errorf("An error occurred. %v", err)
	}
	status := response["status"].(bool)
	if !status {
		t.Errorf("Incorrect response: " + response["message"].(string) )
	}
	payload = []byte(`{"email" : "test6@gmail.com", "password" : "secret", "username" : "landlord1", "role" : 1}`)
	r,_= http.NewRequest("POST", "/api/user/new", bytes.NewBuffer(payload))
	w = httptest.NewRecorder()
	Router().ServeHTTP(w,r)
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.Errorf("An error occurred. %v", err)
	}
	status = response["status"].(bool)
	if !status {
		t.Errorf("Incorrect response: " + response["message"].(string) )
	}
	fmt.Println(response["message"].(string))
}
func TestIncorrectLogin(t *testing.T){
	payload := []byte(`{"email" : "test6@gmail.com", "password" : "wrong"}`)
	r,_:= http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w,r)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.Errorf("An error occurred. %v", err)
	}
	status := response["status"].(bool)
	if status {
		t.Errorf("Incorrect response: " + response["message"].(string) )
	}
	fmt.Println(response["message"].(string))
}
func TestLogin(t *testing.T){
	payload := []byte(`{"email" : "test6@gmail.com", "password" : "secret"}`)
	r,_:= http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w,r)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.Errorf("An error occurred. %v", err)
	}
	status := response["status"].(bool)
	if !status {
		t.Errorf("Incorrect response: " + response["message"].(string) )
	}
	fmt.Println(response["message"].(string))
}
func TestCreatePropertyNoLogin(t *testing.T) {
	payload := []byte(`{
			"room_size":100,
			"rooms":1,
			"bathrooms":1, 
			"parking":1, 
			"floors":2, 
			"latitude":320.001, 
			"longitude":350.21, 
			"pets":true, 
			"property_type":1, 
			"price":500,
			"private_security":true, 
			"capacity":4
			}`)
	r,_:= http.NewRequest("POST", "/api/properties/new", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w,r)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.Errorf("An error occurred. %v", err)
	}
	status := response["status"].(bool)
	if status {
		t.Errorf("Incorrect response: " + response["message"].(string) )
	}
	fmt.Println(response["message"].(string))
}
func TestLogout(t *testing.T) {
	r,_:= http.NewRequest("POST", "/api/user/logout", nil)
	w := httptest.NewRecorder()
	Router().ServeHTTP(w,r)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.Errorf("An error occurred. %v", err)
	}
	status := response["status"].(bool)
	if !status {
		t.Errorf("Incorrect response: " + response["message"].(string) )
	}
	fmt.Println(response["message"].(string))
}
func TestActiveProperties(t *testing.T) {
	r,_:= http.NewRequest("GET", "/api/properties/getActive", nil)
	w := httptest.NewRecorder()
	Router().ServeHTTP(w,r)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.Errorf("An error occurred. %v", err)
	}
	status := response["status"].(bool)
	if !status {
		t.Errorf("Incorrect response: " + response["message"].(string) )
	}
	fmt.Println(response["message"].(string))
}
