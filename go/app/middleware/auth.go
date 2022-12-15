package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (claims jwt.Claims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	// if claims.ExpiresAt < time.Now().Local().Unix() {
	// 	err = errors.New("token expired")
	// 	return
	// }
	// fmt.Printf("%+v", claims)
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := ValidateToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		// email, ok := claims.(JWTClaim)
		// role := claims.(jwt.MapClaims)["Role"].(string)
		// name := claims.(jwt.MapClaims)["Username"].(string)
		// fmt.Printf("%+v", claims)
		payLoad, ok := claims.(*JWTClaim)

		if !ok {
			fmt.Printf("type error")
		}
		// fmt.Printf("%+v %+v", payLoad, ok)

		r.Header.Set("email", payLoad.Email)
		r.Header.Set("role", payLoad.Role)
		r.Header.Set("userName", payLoad.Username)

		next.ServeHTTP(w, r)
	})
}
