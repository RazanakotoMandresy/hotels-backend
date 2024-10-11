package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)
func getUUUIDVars(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	uuid, exist := vars["uuid"]

	if !exist {
		return "", errors.New("invalid uuid in vars")
	}
	return uuid, nil
}