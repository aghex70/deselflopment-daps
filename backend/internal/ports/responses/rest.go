package responses

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const (
	TODO_STRING            string = "todo/"
	CATEGORY_STRING        string = "category/"
	USER_STRING            string = "user/"
	USER_ACTIVATION_STRING string = "activate/"
)

type APIErrorResponse struct {
	Message string `json:"message"`
}

func ValidateRequest(r *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		if err != io.EOF {
			return err
		}
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

func CheckHttpMethod(status string, w http.ResponseWriter, r *http.Request) error {
	if r.Method != status {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New("method not allowed")
	}
	return nil
}
