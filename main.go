package main

import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
	   ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/spf13/viper"
    "log"

    _ "crud_go_app/docs"

    "crud_go_app/models"
    "crud_go_app/controllers"
    "crud_go_app/middlewares"
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

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
    if err := initConfig(); err != nil {
        log.Fatalf("Error initializing configs: %s", err.Error())
    }
    route := gin.Default()

    // Подключение к базе данных
    models.ConnectDB(models.ConfigDB{
        Host:     viper.GetString("db.host"),
        Port:     viper.GetString("db.port"),
        Username: viper.GetString("db.user"),
        Password: viper.GetString("db.password"),
        DBName:   viper.GetString("db.dbname"),
        SSLMode:  viper.GetString("db.sslmode"),
    })

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
    route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    // Маршруты
    route.GET("/tracks", controllers.GetAllTracks)
    route.POST("/tracks", controllers.CreateTrack)
    route.GET("/tracks/:id", controllers.GetTrack)
    route.PATCH("/tracks/:id", controllers.UpdateTrack)
    route.DELETE("/tracks/:id", controllers.DeleteTrack)

    profile := route.Group("/profile").Use(middlewares.Authz())
    {
      profile.GET("/", controllers.Profile)
    }

    public := route.Group("/public")
    {
      public.POST("/login", controllers.Login)
      public.POST("/signup", controllers.Signup)
    }




    // Запуск сервера
    route.Run()
}

func initConfig() error {
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}
