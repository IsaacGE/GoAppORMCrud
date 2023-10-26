package controllers

import (
	response "GoCrudORM/controllers/response"
	templateHandler "GoCrudORM/helpers"
	"GoCrudORM/router"
	"GoCrudORM/types"
	"encoding/json"
	"errors"
)

func CreateUser(response *response.Response, request *router.Request) {
	status, result := GenerateResponse(request)
	response.SendData(status, result)
}

func GenerateResponse(request *router.Request) (int, string) {
	if len(request.Body) == 0 {
		err := errors.New("The request body is missing required data")
		return 400, templateHandler.GetErrorViewTemplate(err)
	}

	var userData types.User
	if err := json.Unmarshal([]byte(request.Body[0]), &userData); err != nil {
		return 400, templateHandler.GetErrorViewTemplate(err)
	}

	htmlContent := templateHandler.GetUserCreatedViewTemplate(userData)
	return 200, htmlContent
}
