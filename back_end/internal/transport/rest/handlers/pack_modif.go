package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Создание новой записи Наборы модификаций
// @Desription Создает новую запись в справочнике
// @Tags Наборы модификаций (pack_modif)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.PackModif
// @Failure 404 {object} mw.HTTPError
// @Router /pack_modif [post]
func CreatePackModif(ctx *gin.Context) {
	var newPackModif models.PackModif
	if err := ctx.BindJSON(&newPackModif); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreatePackModif(newPackModif)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newPackModif)
	}
}

// @Summary Возвращает список PackModif
// @Description Возвращает массив всех pack_modifs
// @Tags PackModif (pack_modifs)
// @Router /pack_modif [get]
func ReadPackModifs(ctx *gin.Context) {
	data, err := services.ReadPackModifs()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

// @Summary Создание новой записи PackModifSp
// @Desription Создает новую запись в справочнике
// @Tags PackModifSp (pack_modif_sp)
// @Router /pack_modif_sp [post]
func CreatePackModifSp(ctx *gin.Context) {
	var newPackModifSp models.PackModifSp
	if err := ctx.BindJSON(&newPackModifSp); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreatePackModifSp(newPackModifSp)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newPackModifSp)
	}
}

// @Summary Возвращает список PackModifSp
// @Description Возвращает массив всех pack_modif_sp
// @Tags Имущество (pack_modif_sp)
// @Router /pack_modif_sp [get]
func ReadPackModifSps(ctx *gin.Context) {
	data, err := services.ReadPackModifSps()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}
