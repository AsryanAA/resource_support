package handlers

import (
	"back/internal/transport/rest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Возвращает рассчет потребности
// @Description Возвращает рассчет потребности по лицевой карточке
// @Tags Calculations
// @Router /calculation [get]
func RequirementCalculations(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := services.RequirementCalculations(int64(id), "2024-02-15")
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
