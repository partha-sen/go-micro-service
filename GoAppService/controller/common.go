package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"learning.go/appservice/model"
)

func writeJson(w http.ResponseWriter, obj interface{}) {

	data, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func parsePathParam(strId string) (int, error) {
	id, err := strconv.Atoi(strId[1:])
	pathErr := errors.Wrap(err, "Could not get from path param")
	return id, pathErr
}

func processGet(w http.ResponseWriter, strId string, m func(id int) (model.Any, error)) {

	id, err := parsePathParam(strId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := m(id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	writeJson(w, value)

}

func processPut(w http.ResponseWriter, obj model.Any, m func(model.Any) (int64, error)) {

	id, err := m(obj)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := []byte(strconv.FormatInt(id, 10))
	w.Write(b)
}

func processGetAll(w http.ResponseWriter, m func() ([]model.Any, error)) {
	values, err := m()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(values) > 0 {
		writeJson(w, values)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func processPost(w http.ResponseWriter, obj model.Any, m func(model.Any) (int64, error)) {

	id, err := m(obj)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := []byte(strconv.FormatInt(id, 10))
	w.Write(b)
}

func processDelete(w http.ResponseWriter, strId string, m func(id int) (int64, error)) {

	id, err := parsePathParam(strId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count, err := m(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := []byte(strconv.FormatInt(count, 10))
	w.Write(b)

}

func requestBodyToObject(w http.ResponseWriter, r *http.Request, p model.Any) error {

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}
