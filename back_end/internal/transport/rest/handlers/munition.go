package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Создание новой записи Имущество
// @Desription Создает новую запись в справочнике
// @Tags Имущество (munitions)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.Munition
// @Failure 404 {object} mw.HTTPError
// @Router /munition [post]
func CreateMunition(ctx *gin.Context) {
	var newMunition models.Munition
	if err := ctx.BindJSON(&newMunition); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreateMunition(newMunition)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newMunition)
	}
}

// @Summary Возвращает список Имущество
// @Description Возвращает массив всех munitions
// @Tags Имущество (munitions)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.Munition
// @Failure 404 {object} mw.HTTPError
// @Router /munition [get]
func ReadMunitions(ctx *gin.Context) {
	data, err := services.ReadMunitions()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

// @Summary Обновление записи Имущество
// @Desription Обноляет запись в справочнике
// @Tags Имущество (munitions)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.Munition
// @Failure 404 {object} mw.HTTPError
// @Router /munition [patch]
func UpdateMunition(ctx *gin.Context) {
	var updateMunition models.Munition
	if err := ctx.BindJSON(&updateMunition); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	err := services.UpdateMunition(updateMunition)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Обноление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, updateMunition.RN)
	}
}

// @Summary Удаление записи Имущество
// @Desription Удаляет запись в справочнике
// @Tags Имущество (munitions)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.Munition
// @Failure 404 {object} mw.HTTPError
// @Router /munition/:id [delete]
func DeleteMunition(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := services.DeleteMunition(int64(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Удаление не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, id)
	}
}