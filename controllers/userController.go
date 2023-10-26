package controllers

import (
	response "GoCrudORM/controllers/response"
	templateHandler "GoCrudORM/helpers"
	"GoCrudORM/router"
	"GoCrudORM/types"
	"errors"
	"fmt"
	"strconv"
)

func CreateUser(response *response.Response, request *router.Request) {
	fmt.Print("REQUEST: ", request)
	status, result := GenerateResponse(request)
	response.SendData(status, result)
}

func GenerateResponse(request *router.Request) (int, string) {
	if request.Body == nil {
		err := errors.New("The request body is nil and it is required")
		return 400, templateHandler.GetErrorViewTemplate(err)
	}

	age, err := strconv.ParseInt(request.Body[1], 10, 64)
	if err != nil {
		return 500, templateHandler.GetErrorViewTemplate(err)
	}

	userData := types.User{
		Name:  request.Body[0],
		Age:   int(age),
		Email: request.Body[3],
	}

	htmlContent := templateHandler.GetUserCreatedViewTemplate(userData)
	return 200, htmlContent
}
