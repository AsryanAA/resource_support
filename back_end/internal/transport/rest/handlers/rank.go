package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateRank(ctx *gin.Context) {
	var newRank models.Rank
	if err := ctx.BindJSON(&newRank); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreateRank(newRank)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newRank)
	}
}

func ReadRanks(ctx *gin.Context) {
	data, err := services.ReadRanks()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

func UpdateRank(ctx *gin.Context) {
	var updateRank models.Rank
	if err := ctx.BindJSON(&updateRank); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	err := services.UpdateRank(updateRank)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Обноление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, updateRank.Id)
	}
}

func DeleteRank(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := services.DeleteRank(int64(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Удаление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, id)
	}
}
