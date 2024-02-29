package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateMunitionMod(ctx *gin.Context) {
	var newMunitionMod models.MunitionMod
	if err := ctx.BindJSON(&newMunitionMod); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreateMunitionMod(newMunitionMod)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newMunitionMod)
	}
}

func ReadMunitionsMod(ctx *gin.Context) {
	data, err := services.ReadMunitionsMod()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

func UpdateMunitionMod(ctx *gin.Context) {
	var updateMunitionMod models.MunitionMod
	if err := ctx.BindJSON(&updateMunitionMod); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	err := services.UpdateMunitionMod(updateMunitionMod)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Обноление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, updateMunitionMod.RN)
	}
}

func DeleteMunitionMod(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := services.DeleteMunitionMod(int64(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Удаление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, id)
	}
}
