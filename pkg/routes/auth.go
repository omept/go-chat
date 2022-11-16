package routes

import (
	"github.com/gorilla/mux"
	"github.com/ong-gtp/go-chat/pkg/controllers"
	"github.com/ong-gtp/go-chat/pkg/domain/middlewares"
)

var RegisterAuthRoutes = func(router *mux.Router) {

	sb := router.PathPrefix("/v1/api/auth").Subrouter()
	sb.Use(middlewares.HeaderMiddleware)

	sb.HandleFunc("/login", controllers.Login).Methods("POST")
	sb.HandleFunc("/signup", controllers.SignUp).Methods("POST")
}
