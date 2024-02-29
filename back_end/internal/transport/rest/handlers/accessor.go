package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Создание новой записи Фурнитура
// @Desription Создает новую запись в справочнике
// @Tags Фурнитура (accessors)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.Accessor
// @Failure 404 {object} mw.HTTPError
// @Router /accessor [post]
func CreateAccessor(ctx *gin.Context) {
	var newAccessor models.Accessor
	if err := ctx.BindJSON(&newAccessor); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	err := services.CreateAccessor(newAccessor)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
		return
	} else {
		ctx.IndentedJSON(http.StatusCreated, newAccessor)
		return
	}
}

// @Summary Возвращает список Фурнитура
// @Description Возвращает массив всех accessors
// @Tags Фурнитура (accessors)
// @Param auth_id body string true "AuthId"
// @Param password_web body string true "PasswordWeb"
// @Success 200 {object} models.Accessor
// @Failure 404 {object} mw.HTTPError
// @Router /accessor [get]
func ReadAccessors(ctx *gin.Context) {
	data, err := services.ReadAccessors()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Чтение не произошло",
		})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
		return
	}
}

// @Summary Обновление записи Фурнитура
// @Desription Обноляет запись в справочнике
// @Tags Фурнитура (accessors)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.Accessor
// @Failure 404 {object} mw.HTTPError
// @Router /accessor [patch]
func UpdateAccessor(ctx *gin.Context) {
	var updateAccessor models.Accessor
	if err := ctx.BindJSON(&updateAccessor); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	err := services.UpdateAccessor(updateAccessor)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Обноление не произошло",
		})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, updateAccessor.RN)
		return
	}
}

// @Summary Удаление записи Фурнитура
// @Desription Удаляет запись в справочнике
// @Tags Фурнитура (accessors)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.Accessor
// @Failure 404 {object} mw.HTTPError
// @Router /accessor/:id [delete]
func DeleteAccessor(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := services.DeleteAccessor(int64(id))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Удаление не произошло",
		})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, id)
		return
	}
}
