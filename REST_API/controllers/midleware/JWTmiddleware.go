package middleware

import (
	secret "RestApi/.env"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const UserId ctxKey = "user_id"

func JWTmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token := bearerToken[1]
				payload := jwt.MapClaims{}
				jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
					return secret.SecretKey, nil
				})
				value, err := payload.GetSubject()
				if err != nil {
					log.Println(err)
					return
				}
				user_id, err := strconv.Atoi(value)
				if err != nil {
					log.Println(err)
					return
				}
				ctx := context.WithValue(r.Context(), UserId, uint(user_id))
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			}
		} else {
			fmt.Fprintf(w, "An authorization header is required.")
		}
	})
}

// func JwtVerify(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authorizationHeader := r.Header.Get("Authorization")
// 		if authorizationHeader != "" {
// 			bearerToken := strings.Split(authorizationHeader, " ")
// 			if len(bearerToken) == 2 {
// 				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
// 					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 						return nil, fmt.Errorf("There was an error in parsing token.")
// 					}
// 					return secret.SecretKey, nil
// 				})
// 				if err != nil {
// 					fmt.Fprintf(w, err.Error())
// 					return
// 				}
// 				if token.Valid {
// 					log.Println(token)
// 					next.ServeHTTP(w, r)
// 				} else {
// 					fmt.Fprintf(w, "Invalid token provided.")
// 				}
// 			}
// 		} else {
// 			fmt.Fprintf(w, "An authorization header is required.")
// 		}
// 	})
// }
