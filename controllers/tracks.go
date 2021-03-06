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


// GetAllTracks godoc
// @Summary Show a Tracks
// @Description Get all tracks
// @ID get-all-tracks-by
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Track
// @Router /tracks [get]
func GetAllTracks(context *gin.Context) {
   var tracks []models.Track  // срез элементы которого трэки из базы (тип models.Track)
   models.DB.Find(&tracks)

   context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}


// CreateTrack godoc
// @Summary Create a new Tracл
// @Description post request example
// @Accept json
// @Produce plain
// @Param message body models.Track true "Track Info"
// @Success 200 {string} string "success"
// @Failure 500 {string} string "fail"
// @Router /tracks/ [post]
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

// GetTrack godoc
// @Summary Get track by id
// @Description Get track by id
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Track ID"
// @Success 200 {object} models.Track
// @Router /tracks/{id} [get]
func GetTrack(context *gin.Context) {
   // Проверяем имеется ли запись
   var track models.Track
   if err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
      context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
      return
   }

   context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// UpdateTrack godoc
// @Summary Change track
// @Description  Change track
// @ID change-string-by-int
// @Accept  json
// @Produce  json
// @Param message body models.Track true "Track Info"
// @Param id path int true "Track ID"
// @Success 200 {object} models.Track
// @Router /tracks/{id} [patch]
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

// DeleteTrack godoc
// @Summary Delete track by id
// @Description Delete track by id
// @ID delete-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Track ID"
// @Success 204
// @Router /tracks/{id} [delete]
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
