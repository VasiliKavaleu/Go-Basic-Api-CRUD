package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"

    "crud_go_app/models"
)


type CreateTrackInput struct {
   Artist string `json:"artist" binding:"required"`
   Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
   Artist string `json:"artist"`
   Title  string `json:"title"`
}

// GET /tracks
// Получаем список всех треков
func GetAllTracks(context *gin.Context) {
   var tracks []models.Track  // срез элементы которого трэки из базы (тип models.Track)
   models.DB.Find(&tracks)

   context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}


// POST /tracks
// Создание трека
func CreateTrack(context *gin.Context) {

   var input CreateTrackInput


   if err := context.ShouldBindJSON(&input); err != nil {
      context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
   }

   track := models.Track{Artist: input.Artist, Title: input.Title}
   models.DB.Create(&track)

   context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// GET /tracks/:id
// Получение одного трека по ID
func GetTrack(context *gin.Context) {
   // Проверяем имеется ли запись
   var track models.Track
   if err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
      context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
      return
   }

   context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// PATCH /tracks/:id
// Изменения информации
func UpdateTrack(context *gin.Context) {
   // Проверяем имеется ли такая запись перед тем как её менять
   var track models.Track
   if err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
      context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
      return
   }

   var input UpdateTrackInput
   if err := context.ShouldBindJSON(&input); err != nil {
      context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
   }

   models.DB.Model(&track).Update(input)

   context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// DELETE /tracks/:id
// Удаление
func DeleteTrack(context *gin.Context) {
   // Проверяем имеется ли такая запись перед тем как её удалять
   var track models.Track
   if err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
      context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
      return
   }

   models.DB.Delete(&track)

   context.JSON(http.StatusNoContent, gin.H{"Message": "Track deleted successefully."})
}
