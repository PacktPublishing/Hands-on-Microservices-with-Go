package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/utils/appErrors"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/usecases"
	"github.com/gorilla/mux"
)

type Handlers struct {
	GetUserUsecase    usecases.GetUser
	UpdateUserUsecase usecases.UpdateUser
}

func (handler *Handlers) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username, ok := vars["username"]
	username = strings.Trim(username, " ")
	if !ok || username == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Username parameter is required.")
		return
	}
	user, err := handler.GetUserUsecase.GetUser(username)
	if err != nil {
		//Verify if it was a 404
		if err == appErrors.ErrorNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Username does not exist on our records.")
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(userJSON))
}

func (handler *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Username parameter is required.")
		return
	}

	//Body has JSON Object
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	user := &entities.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	//Verify it's same user as username
	if username != user.Username {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Username on Query is different than Username on body,")
		return
	}

	//Verify The Data
	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Email is required,")
		return
	}

	if user.FirstName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "FirstName is required,")
		return
	}

	if user.LastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "FirstName is required,")
		return
	}

	if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Password is required,")
		return
	}

	err = handler.UpdateUserUsecase.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User updated Correctly.")
}
