package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
)

// dummy user data
var users = map[string]string{"user1": "password", "user2": "password"}
// creating a cookie session store
var store = sessions.NewCookieStore([]byte("secret_key"))


func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/", app.getMessageLog)
	mux.Post("/login", app.loginHandler)
	mux.Get("/logout", app.logoutHandler)
	mux.Get("/healthcheck", app.healthcheck)

	return mux
}