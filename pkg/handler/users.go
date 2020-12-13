package handler

import (
	"cook-helper/pkg/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) signUp(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newResponseError(rw, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		newResponseError(rw, http.StatusBadRequest, err.Error())
		return
	}

	err = h.validate.Struct(user)
	if err != nil {
		newResponseError(rw, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newResponseError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	newResponseOk(rw, map[string]interface{}{"id": id})
}
