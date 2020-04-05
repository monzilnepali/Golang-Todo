package middleware

import (
	"fmt"
	"net/http"

	jwt "github.com/monzilnepali/Golang-Todo/services"
)

//AuthMiddleware handler
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("passing through middleware")
		//middleware logic here

		//! extracting token from request header
		token := r.Header["Authorization"]
		fmt.Println("token", token)

		if len(token) != 0 {
			//! verify the token
			//* isverified call next http.handler
			//* otherwise unauthorized response
			tokenString := token[0]
			res, err := jwt.VerifyToken(tokenString)
			fmt.Println("logged In", claims.id)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		}

	})
}
