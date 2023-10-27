package controllers

import (
	context "GoCrudORM/context"
	response "GoCrudORM/controllers/response"
	handlers "GoCrudORM/helpers"
	"GoCrudORM/router"
	"GoCrudORM/types"
	"encoding/json"
	"errors"
	"time"
)

/**
 * @api {post} /user Create User
 * @apiName CreateUser
 * @apiGroup User
 * @apiParam {String} name User name
 * @apiParam {String} email User email
 */
func CreateUser(response *response.Response, request *router.Request) {
	if len(request.Body) == 0 {
		err := errors.New("The request body is missing required data")
		response.SendJSON(handlers.HandleJSONResponse400(nil, "", "", &err))
		return
	}

	var userData types.User
	if err := json.Unmarshal([]byte(request.Body[0]), &userData); err != nil {
		response.SendJSON(handlers.HandleJSONResponse400(nil, "", "", &err))
		return
	}

	dbContext := context.DbContext
	userData.Id = 0
	userData.RegisterDate = time.Now().UTC()

	result := dbContext.Create(&userData)
	if result.Error != nil {
		response.SendJSON(handlers.HandleJSONResponse500("", "", (&result.Error)))
		return
	}

	response.SendJSON(handlers.HandleJSONResponse200(userData, "User created successfully", "User"))
}

func UpdateUser(response *response.Response, request *router.Request) {
	if len(request.Body) == 0 {
		err := errors.New("The request body is missing required data")
		response.SendJSON(handlers.HandleJSONResponse400(nil, "", "", &err))
		return
	}

	var userData types.User
	if err := json.Unmarshal([]byte(request.Body[0]), &userData); err != nil {
		response.SendJSON(handlers.HandleJSONResponse400(nil, "", "", &err))
		return
	}

	dbContext := context.DbContext
	userData.RegisterDate = time.Now().UTC()

	result := dbContext.Updates(userData)
	if result.Error != nil {
		response.SendJSON(handlers.HandleJSONResponse500("", "", (&result.Error)))
		return
	}

	response.SendJSON(handlers.HandleJSONResponse200(userData, "User updated successfully", "User"))
}

func DeleteUser(response *response.Response, request *router.Request) {
	if len(request.Body) == 0 {
		err := errors.New("The request body is missing required data")
		response.SendJSON(handlers.HandleJSONResponse400(nil, "", "", &err))
		return
	}

	var userData types.User
	if err := json.Unmarshal([]byte(request.Body[0]), &userData); err != nil {
		response.SendJSON(handlers.HandleJSONResponse400(nil, "", "", &err))
		return
	}

	dbContext := context.DbContext

	result := dbContext.Where("id = ?", userData.Id).Delete(&types.User{})
	if result.Error != nil {
		response.SendJSON(handlers.HandleJSONResponse500("", "", &result.Error))
		return
	}

	response.SendJSON(handlers.HandleJSONResponse200(userData, "User deleted successfully", "User"))
}

/**
 * @api {get} /users Get all users
 * @apiVersion 1.0.0
 * @apiName GetAllUsers
 * @apiGroup Users
 * @apiSuccess {Array} users List of users
 */
func GetAllusers(response *response.Response, request *router.Request) {
	dbContext := context.DbContext
	var users []types.User
	result := dbContext.Find(&users)

	if result.Error != nil {
		response.SendJSON(handlers.HandleJSONResponse500("", "", (&result.Error)))
		return
	}

	response.SendJSON(handlers.HandleJSONResponse200(users, "Users retrieved successfully", "Users"))
}
