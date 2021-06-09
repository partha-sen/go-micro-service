package controller

import (
	"net/http"
	"strings"

	"learning.go/appservice/model"
	"learning.go/appservice/service"
)

func HandleQuestion(w http.ResponseWriter, r *http.Request) {

	strId := strings.TrimPrefix(r.URL.Path, "/questions")

	switch r.Method {

	case http.MethodGet:
		if len(strId) > 0 {
			processGet(w, strId, service.GetQuestionById)
		} else {
			processGetAll(w, service.GetAllQuestion)
		}

	case http.MethodPut:
		id, err := parsePathParam(strId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var obj model.Question
		err = requestBodyToObject(w, r, &obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		obj.Id = id
		callUpdate := func(obj model.Any) (int64, error) {
			return service.UpdateQuestion(obj.(model.Question))
		}
		processPut(w, obj, callUpdate)

	case http.MethodPost:
		var obj model.Question
		err := requestBodyToObject(w, r, &obj)

		if err == nil {
			callSave := func(obj model.Any) (int64, error) {
				return service.SaveQuestion(obj.(model.Question))
			}
			processPost(w, obj, callSave)
		}

	case http.MethodDelete:
		processDelete(w, strId, service.DeleteQuestion)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func HandleQuestions(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		processGetAll(w, service.GetAllQuestion)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
