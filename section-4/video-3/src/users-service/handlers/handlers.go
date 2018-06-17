package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/users-service/repositories"
	"github.com/gorilla/mux"
)

type Handlers struct {
	Repo *repositories.MySQLUserRepository
}

func (handler *Handlers) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user, err := handler.Repo.GetUserByUsername(username)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}
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
