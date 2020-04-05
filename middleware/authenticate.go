package middleware

import (
	"net/http"

	"github.com/monzilnepali/Golang-Todo/model"

	jwt "github.com/monzilnepali/Golang-Todo/services"
)

//AuthMiddleware handler
// func AuthMiddleware(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("passing through middleware")
// 		//middleware logic here

// 		//! extracting token from request header
// 		token := r.Header["Authorization"]

// 		if len(token) != 0 {
// 			//! verify the token
// 			//* isverified call next http.handler
// 			//* otherwise unauthorized response
// 			tokenString := token[0]
// 			res, err := jwt.VerifyToken(tokenString)
// 			if err != nil {
// 				http.Error(w, http.StatusText(401), http.StatusUnauthorized)

// 			}

// 			fmt.Println("logged In id ", res)
// 			//should not use basic type string as key in context.WithValue
// 			type userCtxKeyType int
// 			const userCtxKey userCtxKeyType = "userId"
// 			ctx := context.WithValue(r.Context(), userCtxKey, res)

// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		} else {
// 			http.Error(w, http.StatusText(401), http.StatusUnauthorized)

// 		}

// 	})
// }

//Auth middleware
func Auth(fn func(http.ResponseWriter, *http.Request, *model.User)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//! extracting token from request header
		token := r.Header["Authorization"]

		if len(token) != 0 {
			//! verify the token
			//* isverified call next http.handler
			//* otherwise unauthorized response
			tokenString := token[0]
			res, err := jwt.VerifyToken(tokenString)
			if err != nil {
				http.Error(w, http.StatusText(401), http.StatusUnauthorized)

			}
			activeUser := model.User{UserID: res}
			fn(w, r, &activeUser)
		} else {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)

		}
	}

}
