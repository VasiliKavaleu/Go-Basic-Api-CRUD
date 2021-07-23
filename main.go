package main

import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

    _ "crud_go_app/docs"

    "crud_go_app/models"
    "crud_go_app/controllers"
)


// @title Basic CRUD App via Golang
// @version 1.0
// @description Swagger API for Golang Project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email example@email.com

// @license.name MIT
// @license.url https://github.com/VasiliKavaleu/Go-Basic-Api-CRUD.git

// @BasePath /

func main() {
    route := gin.Default()

    // Подключение к базе данных
    models.ConnectDB()

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    // Маршруты
    route.GET("/tracks", controllers.GetAllTracks)
    route.POST("/tracks", controllers.CreateTrack)
    route.GET("/tracks/:id", controllers.GetTrack)
    route.PATCH("/tracks/:id", controllers.UpdateTrack)
    route.DELETE("/tracks/:id", controllers.DeleteTrack)


    // Запуск сервера
    route.Run()
}
