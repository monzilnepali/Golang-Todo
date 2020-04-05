package middleware

import (
	"fmt"
	"net/http"
)

//AuthMiddleware handler
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("passing through middleware")
		//middleware logic here
		next.ServeHTTP(w, r)
	})
}
