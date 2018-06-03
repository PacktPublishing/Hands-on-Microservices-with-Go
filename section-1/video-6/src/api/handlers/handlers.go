package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-1/video-6/src/api/repository"
)

type Handlers struct {
	Repo *repository.Repository
}

func (h *Handlers) All(w http.ResponseWriter, r *http.Request) {
	tzcs, err := h.Repo.FindAll()
	if err != nil {
		error500(w, err)
		return
	}
	jr, err := json.Marshal(tzcs)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, string(jr))
}

func (h *Handlers) GetByTZ(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tz, ok := params["timeZone"]
	if !ok {
		error400(w, "timeZone is Required.")
		return
	}
	tzc, err := h.Repo.FindByTimeZone(tz)
	if err != nil {
		error500(w, err)
		return
	}
	jr, err := json.Marshal(tzc)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, string(jr))
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tz, ok := params["timeZone"]
	if !ok {
		error400(w, "timeZone is Required.")
		return
	}

	tzc := repository.TZConvertion{
		TimeZone: tz,
	}

	err := h.Repo.Delete(tzc)
	if err != nil {
		error500(w, err)
		return
	}

	ok200(w, "Element succesfully deleted.")
}

func (h *Handlers) Insert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var tzc repository.TZConvertion
	err := json.NewDecoder(r.Body).Decode(&tzc)
	if err != nil {
		error400(w, "Invalid json.")
		return
	}

	err = h.Repo.Insert(tzc)
	if err != nil {
		error500(w, err)
		return
	}

	ok200(w, "Element succesfully inserted.")
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tz, ok := params["timeZone"]
	if !ok {
		error400(w, "timeZone is Required.")
		return
	}

	var tzc repository.TZConvertion
	err := json.NewDecoder(r.Body).Decode(&tzc)
	if err != nil {
		error400(w, "Invalid json.")
		return
	}

	err = h.Repo.Update(tz, tzc)
	if err != nil {
		error500(w, err)
		return
	}

	ok200(w, "Element succesfully updated.")
}
