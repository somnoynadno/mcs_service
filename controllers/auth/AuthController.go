package auth

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mcs_service/db"
	"mcs_service/models/auxiliary"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
	"os"
	"time"
)

func authenticate(username, password string) map[string]interface{} {
	user := &entities.User{}
	err := db.GetDB().Preload("UserRole").Where("username = ?", username).First(user).Error

	if err != nil {
		log.Warn(err)
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if !u.CheckPasswordHash(password, user.Password) { // Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	// Worked! Logged In
	db.GetDB().Model(&user).Update("LastLogin", time.Now())

	// Create JWT token
	tk := &auxiliary.Token{UserID: user.ID, UserRole: user.UserRole.Name}

	tk.ExpiresAt = time.Now().Add(time.Hour * 72).Unix() // valid for 3 days
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	resp := u.Message(true, "Logged In")
	resp["token"] = tokenString

	return resp
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	account := &auxiliary.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	resp := authenticate(account.Username, account.Password)
	if resp["token"] == nil {
		u.HandleBadRequest(w, errors.New("wrong credentials"))
		return
	}

	u.Respond(w, resp)
}
