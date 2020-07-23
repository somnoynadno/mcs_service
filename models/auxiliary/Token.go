package auxiliary

import "github.com/dgrijalva/jwt-go"

/*
JWT claims struct
*/
type Token struct {
	UserID   uint
	UserRole string
	jwt.StandardClaims
}

