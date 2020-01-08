package main

/*import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "propertiesWeb"
)
var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
	db *sql.DB
)
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getPwd() []byte {
	fmt.Println("Enter a password")
	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates := template.Must(template.ParseFiles("templates/login.html"))
		templates.ExecuteTemplate(w, "login.html", "")
	}else{
		r.ParseForm()
		fmt.Println(r.Form)
		plainPassword := []byte(r.Form["password"][0])
		rows, err := db.Query("SELECT * FROM \"rentWebsite\".users where username = $1",r.Form["username"][0])
		checkErr(err)
		var hashedPwd string
		var id int
		var username string
		var role int
		for rows.Next(){
			err = rows.Scan(&id,&username,&hashedPwd,&role)
			checkErr(err)
		}
		if comparePasswords(hashedPwd,plainPassword) {
			session, _ := store.Get(r, "cookie-name")
			fmt.Println("match!")
			session.Values["authenticated"] = true
			session.Values["userID"] = id
			session.Values["userRole"] = role
			err = session.Save(r, w)
			checkErr(err)
			http.Redirect(w,r,"/home",http.StatusSeeOther)
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func logout(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}
func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		fmt.Println(ok)
		fmt.Println(auth)
		http.Redirect(w,r,"/login",http.StatusSeeOther)
	} else {
		templates := template.Must(template.ParseFiles("templates/index.html"))
		templates.ExecuteTemplate(w, "index.html", "")
	}
}
func properties(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		http.Redirect(w,r,"/login",http.StatusSeeOther)
	} else {
		templates := template.Must(template.ParseFiles("templates/property-grid.html"))
		templates.ExecuteTemplate(w, "property-grid.html", "")
	}

}
func createProperty(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		http.Redirect(w,r,"/login",http.StatusSeeOther)
	} else {
		if r.Method == "POST" {
			r.ParseForm()
			user := session.Values["userID"]
			roomSize := r.Form["roomSize"][0]
			rooms := r.Form["rooms"][0]
			bathrooms := r.Form["bathrooms"][0]
			parking := r.Form["parking"][0]
			floors := r.Form["floors"][0]
			latitude := r.Form["latitude"][0]
			longitude := r.Form["longitude"][0]
			pets := r.Form["pets"][0]
			propertyType := r.Form["propertyType"][0]
			price := r.Form["price"][0]
			privateSecurity := r.Form["privateSecurity"][0]
			capacity := r.Form["capacity"][0]
			photos := "{c/doc}"
			active := r.Form["active"][0]
			rows, err :=db.Query("INSERT INTO \"rentWebsite\".properties(\"user\", \"roomSize\", rooms, bathrooms, parking, floors, latitude, longitude, pets, \"propertyType\", price, \"privateSecurity\", capacity, photos, active)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);",user,roomSize,rooms,bathrooms,parking,floors,latitude,longitude,pets,propertyType,price,privateSecurity,capacity,photos,active)
			checkErr(err)
			fmt.Println(rows)
		}
	}
}
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//pwd := getPwd()
	//hash := hashAndSalt(pwd)
	//fmt.Println("Salted Hash", hash)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home",home)
	http.HandleFunc("/login",login)
	http.HandleFunc("/logout",logout)
	http.HandleFunc("/properties",properties)
	http.HandleFunc("/createProperty",createProperty)
	fmt.Println(http.ListenAndServe(":8080", nil));
}*/
import (
	"./controllers"
	_ "./models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
func main() {
	router := mux.NewRouter()
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
	//router.Use(app.JwtAuthentication)
	fmt.Println(http.ListenAndServe(":8080", router));

}