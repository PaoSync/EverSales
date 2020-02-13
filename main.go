package main

import (
	_ "./models"
	"net/http"
)
func main() {
	http.ListenAndServe(":8080", Router())
	/*router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/logout", controllers.Logout).Methods("POST")
	router.HandleFunc("/api/properties/getActive", controllers.ActiveProperties).Methods("GET")
	router.HandleFunc("/api/properties/new", controllers.CreateProperty).Methods("POST")
	router.HandleFunc("/api/properties/delete/{id}",controllers.DeletePropertyByID).Methods("POST")
	router.HandleFunc("/api/properties/toggleStatus/{id}",controllers.TogglePropertyStatus).Methods("POST")
	router.HandleFunc("/api/properties/propertyInformation/{id}",controllers.PropertyInformation).Methods("GET")
	router.HandleFunc("/api/properties/modify/{id}",controllers.ModifyProperty).Methods("POST")
	router.HandleFunc("/api/properties/search",controllers.SearchProperties).Methods("POST")
	router.HandleFunc("/api/visits/new/{propertyID}",controllers.NewVisit).Methods("POST")
	fmt.Println(http.ListenAndServe(":8080", router));*/

}

