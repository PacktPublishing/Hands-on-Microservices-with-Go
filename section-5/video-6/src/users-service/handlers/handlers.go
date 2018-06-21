package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-6/src/users-service/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-6/src/users-service/usecases"
	"github.com/gorilla/mux"
)

type Handlers struct {
	GetUserUsecase    *usecases.GetUserUsecase
	UpdateUserUsecase *usecases.UpdateUserUsecase
}

func (handler *Handlers) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user, err := handler.GetUserUsecase.GetUser(username)
	//VERIFICAR TYPO DE ERROR - 404
	if err != nil {
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

func (handler *Handlers) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
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
		return
	}

	//I SHOULD VERIFY THE DATA
	//......

	err = handler.UpdateUserUsecase.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User updated Correctly.")
}

/*
func (handler *Handlers) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userIDstr, ok := vars["userID"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	user, err := handler.Repo.GetUserByID(uint32(userID))
	if err != nil {
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
*/
