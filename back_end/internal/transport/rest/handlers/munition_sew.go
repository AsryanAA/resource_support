package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateMunitionSew(ctx *gin.Context) {
	var newMunitionSew models.MunitionSew
	if err := ctx.BindJSON(&newMunitionSew); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreateMunitionSew(newMunitionSew)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newMunitionSew)
	}
}

func ReadMunitionsSew(ctx *gin.Context) {
	data, err := services.ReadMunitionsSew()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

func DeleteMunitionSew(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := services.DeleteMunitionSew(int64(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Удаление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, id)
	}
}

func UpdateMunitionSew(ctx *gin.Context) {
	var updateMunitionSew models.MunitionSew
	if err := ctx.BindJSON(&updateMunitionSew); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	err := services.UpdateMunitionSew(updateMunitionSew)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Обноление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, updateMunitionSew.RN)
	}
}
