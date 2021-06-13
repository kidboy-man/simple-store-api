package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-store-api/datatransfers"
)

type JSONResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"` // http status code
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
	// TODO pagination
}

func doReturnOK(obj interface{}) (response JSONResponse) {
	response.Success = true
	response.Status = http.StatusOK
	json, err := json.Marshal(obj)
	if err != nil {
		log.Panic(err)
	}
	response.Data = string(json)
	return
}

func doReturnNotOK(err error) (response JSONResponse) {
	response.Success = false
	response.Status = http.StatusInternalServerError
	if v, ok := err.(*datatransfers.CustomError); ok {
		response.Error = v
		response.Status = v.Status
		return
	}

	response.Error = err
	return
}

func ReturnJSONResponse(obj interface{}, err error) (response JSONResponse) {
	if err != nil {
		return doReturnNotOK(err)
	}

	return doReturnOK(obj)
}
