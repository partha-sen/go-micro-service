package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

func parsePathParam(strId string) (int, error) {
	id, err := strconv.Atoi(strId)
	pathErr := errors.Wrap(err, "Could not get from path param")
	return id, pathErr
}

func requestBodyToObject(w http.ResponseWriter, r *http.Request, any interface{}) error {

	err := json.NewDecoder(r.Body).Decode(any)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}
