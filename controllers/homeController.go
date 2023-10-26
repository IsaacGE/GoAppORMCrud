package controllers

import (
	response "GoCrudORM/controllers/response"
	templateHandler "GoCrudORM/helpers"
	"GoCrudORM/router"
	"fmt"
)

func HomeView(response *response.Response, request *router.Request) {
	fmt.Printf("%v\t %v\n", request.Method, request.Route)
	response.SendData(200, GenerateHTMLResponse())
}

func GenerateHTMLResponse() string {
	htmlContent := templateHandler.GetHomeViewTemplate()

	return htmlContent
}
