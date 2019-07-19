package router

import (
	"github.com/gorilla/mux"
	authApi "mudiralmaham/api/authentication"
	"mudiralmaham/api/project"
	"mudiralmaham/api/task"
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
	projectHandler := Router.PathPrefix("/project").Subrouter()
	taskHandler := Router.PathPrefix("/task").Subrouter()
	auth.HandleFunc("/login", authApi.Login)
	auth.HandleFunc("/signUp", authApi.SignUp)
	//auth.HandleFunc("/say_hello", authApi.SayHello)
	auth.Handle("/say_hello", jwtMiddleWare.Handler(http.HandlerFunc(authApi.SayHello))).Methods("GET")
	projectHandler.Handle("/add", jwtMiddleWare.Handler(http.HandlerFunc(project.Add))).Methods("POST")
	projectHandler.Handle("/get", jwtMiddleWare.Handler(http.HandlerFunc(project.Get))).Methods("POST")
	projectHandler.Handle("/add_collaborator", jwtMiddleWare.Handler(http.HandlerFunc(project.AddCollaborator))).Methods("POST")

	taskHandler.Handle("/add", jwtMiddleWare.Handler(http.HandlerFunc(task.Add))).Methods("POST")
	taskHandler.Handle("/get", jwtMiddleWare.Handler(http.HandlerFunc(task.Get))).Methods("POST")
	taskHandler.Handle("/update", jwtMiddleWare.Handler(http.HandlerFunc(task.Update))).Methods("POST")
}
