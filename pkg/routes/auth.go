package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ong-gtp/go-chat/pkg/controllers"
)

var RegisterAuthRoutes = func(router *mux.Router) {

	sb := router.PathPrefix("/v1/api/auth/").Subrouter()
	sb.Use(headerMiddleware)

	sb.HandleFunc("/login/", controllers.Login).Methods("POST")
	sb.HandleFunc("/signup/", controllers.SignUp).Methods("POST")
}

var headerMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
