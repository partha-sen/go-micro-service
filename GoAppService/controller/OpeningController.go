package controller

import (
	"log"
	"net/http"
	"strings"

	"learning.go/appservice/model"
	"learning.go/appservice/service"
	"learning.go/appservice/token"
)

func HandleOpening(w http.ResponseWriter, r *http.Request) {

	strId := strings.TrimPrefix(r.URL.Path, "/openings")

	jwtClaim := r.Context().Value(token.JWT_KEY).(model.JwtClaim)

	if err := token.IsLoggedOut(jwtClaim); err != nil {
		log.Println("IsLoggedOut ", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	switch r.Method {

	case http.MethodGet:
		if len(strId) > 0 {
			processGet(w, strId, service.GetOpeningById)
		} else {
			processGetAll(w, service.GetAllOpening)
		}

	case http.MethodPut:
		id, err := parsePathParam(strId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var obj model.Opening
		err = requestBodyToObject(w, r, &obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		obj.Id = id
		callUpdate := func(obj model.Any) (int64, error) {
			return service.UpdateOpening(obj.(model.Opening))
		}
		processPut(w, obj, callUpdate)

	case http.MethodPost:
		var obj model.Opening
		err := requestBodyToObject(w, r, &obj)

		if err == nil {
			callSave := func(obj model.Any) (int64, error) {
				return service.SaveOpening(obj.(model.Opening))
			}
			processPost(w, obj, callSave)
		}

	case http.MethodDelete:
		processDelete(w, strId, service.DeleteOpening)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
