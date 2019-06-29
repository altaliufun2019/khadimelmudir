package router

import (
	"github.com/gorilla/mux"
	authApi "mudiralmaham/api/authentication"
	"mudiralmaham/api/event"
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
	eventHandler := Router.PathPrefix("/event").Subrouter()
	auth.HandleFunc("/login", authApi.Login)
	auth.HandleFunc("/signUp", authApi.SignUp)
	//auth.HandleFunc("/say_hello", authApi.SayHello)
	auth.Handle("/say_hello", jwtMiddleWare.Handler(http.HandlerFunc(authApi.SayHello))).Methods("GET")
	eventHandler.Handle("/add", jwtMiddleWare.Handler(http.HandlerFunc(event.Add))).Methods("POST")
	eventHandler.Handle("/get", jwtMiddleWare.Handler(http.HandlerFunc(event.Get))).Methods("POST")
}
