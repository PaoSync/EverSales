package models

import (
	u "../utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}
type Account struct {
	gorm.Model
	Username string `json:"username"`
	Email	 string `json:"email"`
	Password string `json:"password"`
	Role     int `json:"role"`
	Token    string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}
	err = GetDB().Table("accounts").Where("username = ?", account.Username).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Username != "" {
		return u.Message(false, "Username already in use."), false
	}
	return u.Message(false, "Requirement passed"), true
}
func (account *Account) Create() (map[string]interface{}) {
	if resp, ok := account.Validate(); !ok {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	GetDB().Create(account)
	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}
	account.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}
func Login(email, password string) (map[string]interface{}){
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?",email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	account.Password = ""
	resp := u.Message(true,"Logged In")
	resp["account"] = account
	return resp
}
