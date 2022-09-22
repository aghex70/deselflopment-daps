package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type APIErrorResponse struct {
	Message string `json:"message"`
}

type MalformedRequest struct {
	Status  int
	Message string
}

func (mr *MalformedRequest) Error() string {
	return mr.Message
}

func ValidateRequest(r *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Printf("%+v", err)
		fmt.Println(err.Error())
		return err
	}

	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		fmt.Printf("%+v", err)
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func ThrowError(err error, status int, w http.ResponseWriter) {
	resp := APIErrorResponse{
		Message: err.Error(),
	}
	data, _ := json.Marshal(resp)
	fmt.Printf("err %+v", err.Error())
	w.WriteHeader(status)
	w.Write(data)
}
