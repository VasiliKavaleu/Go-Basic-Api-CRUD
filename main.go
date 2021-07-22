package main

import (
    "github.com/gin-gonic/gin"

    "crud_go_app/models"
    "crud_go_app/controllers"
)

func main() {
    route := gin.Default()

    // Подключение к базе данных
    models.ConnectDB()

    // Маршруты
    route.GET("/tracks", controllers.GetAllTracks)
    route.POST("/tracks", controllers.CreateTrack)
    route.GET("/tracks/:id", controllers.GetTrack)
    route.PATCH("/tracks/:id", controllers.UpdateTrack)
    route.DELETE("/tracks/:id", controllers.DeleteTrack)


    // Запуск сервера
    route.Run()
}
