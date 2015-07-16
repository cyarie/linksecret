package handlers

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/cyarie/linksecret/application/environment"
)

// Rolling up a custom handler to better deal with HTTP errors and the such. Also lets us pass through an environment
// struct/variables without having to use globals.
type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type WebHandler struct {
	*environment.Env
	H func(e *environment.Env, w http.ResponseWriter, r *http.Request) error
}

// Write a function receiver to integrate the structs we made up above
func (h WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
			case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// INDEX HANDLER
func IndexHandler(env *environment.Env, w http.ResponseWriter, r *http.Request) error {
	var err error

	fmt.Fprintf(w, "WELCOME TO GORT")

	return err
}

// GENERATION HANDLER
func GenerateLink(env *environment.Env, w http.ResponseWriter, r *http.Request) error {
	var err error

	return err
}

// REDIRECT HANDLER
func RedirectHash(env *environment.Env, w http.ResponseWriter, r *http.Request) error {
	var err error
	rc := env.CL

	vars := mux.Vars(r)

	hl := vars["linkHash"]
	rl, err := rc.HGet(hl, "link").Result()
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, rl, 301)

	return err
}
