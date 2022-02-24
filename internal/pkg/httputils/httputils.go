/*
 * @Author: Adrian Faisal
 * @Date: 04/09/21 1.42 PM
 */

package httputils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")

	byteData, err := json.Marshal(data)
	if err != nil {
		log.Printf("[httputils.RespondWithJSON] error unmarshal json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(byteData)
	if err != nil {
		log.Printf("[httputils.RespondWithJSON] write data failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type responseErr struct {
	Error      error `json:"error"`
	StatusCode int   `json:"status_code"`
}

func RespondErrWithJSON(w http.ResponseWriter, statusCode int, err error) {
	resp := responseErr{
		Error:      err,
		StatusCode: statusCode,
	}

	RespondWithJSON(w, statusCode, resp)
}
