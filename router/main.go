package router

import (
	"student-management-system/app/middlewares"
	"student-management-system/handlers"
	"student-management-system/handlers/auth"
	"student-management-system/handlers/users"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(appRoute *fiber.App) {

	
	appRoute.Get("/health", handlers.HandleHealthCheck)

	api := appRoute.Group("/api")

	// database seeder
	api.Options("/seed", handlers.SeedData)

	//auth routes
	api.Post("/login", auth.Login)
	
	//middlware
	api.Use(middlewares.JwtMiddleware)

	// setup the todos group
	todos := api.Group("/todos")
	todos.Get("/",handlers.HandleAllTodos)
	todos.Post("/", handlers.HandleCreateTodo)
	todos.Put("/:id", handlers.HandleUpdateTodo)
	todos.Get("/:id", handlers.HandleGetOneTodo)
	todos.Delete("/:id", handlers.HandleDeleteTodo)
	
	
	// setup the todos group
	user := api.Group("/users")
	user.Get("/",users.GetAllUsers)
	user.Post("/create",users.CreateUser)
}
