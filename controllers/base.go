package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool         `json:"success"`
	Status  int          `json:"status"` // http status code
	Data    interface{}  `json:"data"`
	Error   *CustomError `json:"error"`
	// TODO pagination
}

type CustomError struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"` // http status code
	Message string `json:"message"`
}

// type Paging struct {
// 	TotalPage int `json:"totalPage"`
// 	CurrentPage int `json:"currentPage"`

// }

func ReturnOK(obj interface{}) (response Response) {
	response.Success = true
	response.Status = http.StatusOK
	json, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	response.Data = string(json)
	return
}

func ReturnError(customError *CustomError) (response Response) {
	response.Error = customError
	return
}
