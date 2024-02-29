package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Создание новой записи Имущество
// @Desription Создает новую запись в справочнике
// @Tags Имущество (norm_munits)
// @Router /norm_munit [post]
func CreateNormMunit(ctx *gin.Context) {
	var newNormMunit models.NormMunit
	if err := ctx.BindJSON(&newNormMunit); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreateNormMunit(newNormMunit)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newNormMunit)
	}
}

// @Summary Возвращает список Имущество
// @Description Возвращает массив всех norm_munits
// @Tags Имущество (norm_munits)
// @Router /norm_munit [get]
func ReadNormMunit(ctx *gin.Context) {
	data, err := services.ReadNormMunits()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

// @Summary Создание новой записи Имущество
// @Desription Создает новую запись в справочнике
// @Tags Имущество (norm_munits_sp)
// @Router /norm_munit_sp [post]
func CreateNormMunitSp(ctx *gin.Context) {
	var newNormMunitSp models.NormMunitSp
	if err := ctx.BindJSON(&newNormMunitSp); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	err := services.CreateNormMunitSp(newNormMunitSp)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Создание не произошло",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, newNormMunitSp)
	}
}

// @Summary Возвращает список Имущество
// @Description Возвращает массив всех norm_munits_sp
// @Tags Имущество (norm_munits_sp)
// @Router /norm_munit_sp [get]
func ReadNormMunitSp(ctx *gin.Context) {
	data, err := services.ReadNormMunitsSp()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}

// @Summary Возвращает список Имущество
// @Description Возвращает массив всех norm_munits_sp_alt
// @Tags Имущество (norm_munits_sp_alt)
// @Router /norm_munit_sp_alt [get]
func ReadNormMunitSpAlt(ctx *gin.Context) {
	data, err := services.ReadNormMunitsSpAlt()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Чтение не произошло"})
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}
}
