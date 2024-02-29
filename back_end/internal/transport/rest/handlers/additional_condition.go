package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAdditionalCondition(ctx *gin.Context) {
	var newAdditionalCondition models.AdditionalCondition
	if err := ctx.BindJSON(&newAdditionalCondition); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreateAdditionalCondition(newAdditionalCondition)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newAdditionalCondition)
	}
}

func ReadAdditionalConditions(ctx *gin.Context) {
	data, err := services.ReadAdditionalConditions()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

func UpdateAdditionalCondition(ctx *gin.Context) {
	var updateAdditionalCondition models.AdditionalCondition
	if err := ctx.BindJSON(&updateAdditionalCondition); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	err := services.UpdateAdditionalCondition(updateAdditionalCondition)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Обноление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, updateAdditionalCondition.Id)
	}
}

func DeleteAdditionalCondition(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := services.DeleteAdditionalCondition(int64(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Удаление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, id)
	}
}
