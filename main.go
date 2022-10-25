package main

import (
	"Project1/docs"
	"Project1/todos"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Project 1 Kampus Merdeka X Hacktiv8 ToDos
// @version 1.0
// @description This is a sample server celler server.
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @query.collection.format multi
func main() {
	router := gin.Default()

	router.GET("/todos", todos.GetTodos)
	router.GET("/todos/:id", todos.GetTodosByID)
	router.POST("/todos", todos.AddTodos)
	router.PUT("/todos/:id", todos.UpdateTodos)
	router.DELETE("/todos/:id", todos.DeleteTodos)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL("127.0.0.1:8000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":8000")
}
