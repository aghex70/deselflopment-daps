package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type APIErrorResponse struct {
	Message string `json:"message"`
}

func ValidateRequest(r *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		return err
	}

	return nil
}

func ThrowError(err error, status int, w http.ResponseWriter) {
	resp := APIErrorResponse{
		Message: err.Error(),
	}
	data, _ := json.Marshal(resp)
	w.WriteHeader(status)
	_, werr := w.Write(data)
	if werr != nil {
		return
	}
}
