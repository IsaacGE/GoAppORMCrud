package handlers

import (
	"GoCrudORM/types"
)

func HandleJSONResponse200(data interface{}, msg string, title string) types.ResponseApi {
	return types.ResponseApi{
		Status:  200,
		Message: msg,
		Data:    data,
		Error:   nil,
		Title:   title,
	}
}

func HandleJSONResponse400(data *interface{}, msg string, title string, err *error) types.ResponseApi {
	return types.ResponseApi{
		Status:  400,
		Message: msg,
		Data:    data,
		Error:   *err,
		Title:   title,
	}
}

func HandleJSONResponse500(msg string, title string, err *error) types.ResponseApi {
	return types.ResponseApi{
		Status:  400,
		Message: msg,
		Data:    nil,
		Error:   *err,
		Title:   title,
	}
}
