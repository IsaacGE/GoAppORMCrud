package router

import (
	controller "GoCrudORM/controllers"
	"GoCrudORM/router"
)

func AddRoutes(servevrRouter *router.Router) {
	servevrRouter.AddRoute("/home", controller.HomeView)
	servevrRouter.AddRoute("/createUser", controller.CreateUser)
	servevrRouter.AddRoute("/getAllUsers", controller.GetAllusers)
	servevrRouter.AddRoute("/deleteUser", controller.DeleteUser)
	servevrRouter.AddRoute("/updateUser", controller.UpdateUser)
}
