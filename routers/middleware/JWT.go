package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"mcs_service/models/auxiliary"
	u "mcs_service/utils"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetTokenFromHeader(r *http.Request) (*auxiliary.Token, error) {
	tokenHeader := r.Header.Get("Authorization") // grab the token from the header

	if tokenHeader == "" {
		return nil, errors.New("token is missing")
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return nil, errors.New("invalid/malformed auth token: " + tokenHeader)
	}

	tokenPart := splitted[1] // grab the token part, what we are truly interested in
	tk := &auxiliary.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	if err != nil { // malformed token, returns with http code 403 as usual
		return nil, errors.New("malformed authentication token: " + err.Error())
	}

	if !token.Valid { // token is invalid, maybe not signed on this server
		return nil, errors.New("token is not valid")
	}

	return tk, nil
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tk, err := GetTokenFromHeader(r)
		if err != nil {
			u.HandleUnauthorized(w, errors.New("missing auth token"))
			return
		}

		if tk.ExpiresAt < time.Now().Unix() { // token is expired
			u.HandleUnauthorized(w, errors.New("token is expired"))
			return
		}

		// everything is fine, let's create the context
		v := u.Values{M: map[string]string {
			"user_id": fmt.Sprint(tk.UserID),
			"user_role": tk.UserRole,
		}}

		ctx := context.WithValue(r.Context(), "context", v)
		r = r.WithContext(ctx)

		// useful for monitoring
		log.Debug("user id: ", tk.UserID, ", user role: ", tk.UserRole)

		next.ServeHTTP(w, r)
	})
}
