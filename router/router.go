package router

import (
	"github.com/gorilla/mux"
	authentication2 "mudiralmaham/api/authentication"
	"mudiralmaham/utils/authentication"
	"net/http"
)

var (
	Router        = mux.NewRouter()
	jwtMiddleWare = authentication.JwtMiddleware
)

/**
#main api router
	it maps incoming routes to dedicated functions
*/
func ApiMapper() {
	auth := Router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", authentication2.Login)
	auth.HandleFunc("/signUp", authentication2.SignUp)
	auth.Handle("/say_hello", jwtMiddleWare.Handler(http.HandlerFunc(authentication2.SayHello))).Methods("GET")

}
