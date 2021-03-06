package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	jwt "github.com/monzilnepali/Golang-Todo/services"
)

// RequestData hold json data of useid and todoid
type RequestData struct {
	UserID int `json:"userid"`
	TodoID int `json:"todoid"`
}

//Auth middleware
func Auth(h httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		fmt.Println("middleware ")
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
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "userID", res)
			r = r.WithContext(ctx)
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)
			return
		}

	}

}
