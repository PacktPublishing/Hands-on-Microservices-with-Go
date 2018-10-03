package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetAllUserVideos(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userID := ctx.Value(UserID).(uint32)

	allUserVideos, err := h.GetAllUserVideosUC.GetAllVideosFromUser(userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(allUserVideos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))

}
