package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	utility "go-contacts/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

//a struct to rep user account
type Account struct {
	gorm.Model // create "id" , "created_at" , "updated_at" , "deleted_at"  in database
	//Only exported fields will be encoded/decoded in JSON
	Email    string `json:"email";sql:"not null;unique;type:varchar(100)"` // Set field as not nullable and unique`
	Password string `json:"password";sql:"not null;type:varchar(100)`
	Token    string `json:"token";sql:"-"` // sql:"-" ==> don't save in db
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return utility.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return utility.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	// gorm First() function fill  temp variavble by first object in the table , if first() function has error
	//get it in Error
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utility.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utility.Message(false, "Email address already in use by another user."), false
	}

	return utility.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string]interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account) // add account to the table

	if account.ID <= 0 {
		return utility.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	token_expire_seconds ,_ := utility.StringToInt(os.Getenv("token_expire_seconds"))
	expireToken := time.Now().Add(time.Second * time.Duration( token_expire_seconds ) ).Unix()
	tk := &Token{UserId: account.ID , Email: account.Email , StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireToken,
		Issuer:    "bimclouder-auth",
	}}
	//token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	token := jwt.NewWithClaims( jwt.SigningMethodHS256 , tk )
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	//ttl := 60 * time.Second
	//myClaims["exp"] = time.Now().UTC().Add(ttl).Unix()


	account.Password = "" //delete password

	response := utility.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) (map[string]interface{}) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utility.Message(false, "Email address not found")
		}
		return utility.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return utility.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	token_expire_seconds ,_ := utility.StringToInt(os.Getenv("token_expire_seconds"))
	expireToken := time.Now().Add(time.Second * time.Duration( token_expire_seconds ) ).Unix()
	tk := &Token{UserId: account.ID , Email: account.Email , StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireToken,
		Issuer:    "bimclouder-auth",
	}}
	//token := jwt.NewWithClaims( jwt.GetSigningMethod("HS256"), tk )
	token := jwt.NewWithClaims( jwt.SigningMethodHS256 , tk )
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := utility.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
