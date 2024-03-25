package rest

import (
	docs "back/api"
	"back/internal/transport/rest/handlers"
	"back/internal/transport/rest/mw"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(host string, port int) error {
	router := gin.Default()

	// Обработчик OPTIONS запросов
	//router.Use(func(ctx *gin.Context) {
	//	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	//	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	//	if ctx.Request.Method == "OPTIONS" {
	//		ctx.AbortWithStatus(200)
	//		return
	//	}
	//	ctx.Next()
	//})

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	auth := router.Group("auth")
	{
		auth.POST("login", handlers.SingIn)
	}

	api := router.Group("api")
	{
		// api.Use(mw.CheckToken())
		v1 := api.Group("v1")
		{
			calculations := v1.Group("calculations")
			{
				calculations.GET("base/:id", handlers.RequirementCalculations)
				calculations.GET("actual/:id/:calc_date", handlers.BaseCalculations)
			}

			position := v1.Group("position")
			{
				position.POST("create", handlers.CreatePosition)
				position.GET("read", handlers.ReadPositions)
				position.PATCH("update", handlers.UpdatePosition)
				position.DELETE("delete/:id", handlers.DeletePosition)
			}

			division := v1.Group("division")
			{
				division.POST("create", handlers.CreateDivision)
				division.GET("read", handlers.ReadDivisions)
				division.PATCH("update", handlers.UpdateDivision)
				division.DELETE("delete/:id", handlers.DeleteDivision)
			}

			rank := v1.Group("rank")
			{
				rank.POST("create", handlers.CreateRank)
				rank.GET("read", handlers.ReadRanks)
				rank.PATCH("update", handlers.UpdateRank)
				rank.DELETE("delete/:id", handlers.DeleteRank)
			}

			addCond := v1.Group("add_cond")
			{
				addCond.POST("create", handlers.CreateAdditionalCondition)
				addCond.GET("read", handlers.ReadAdditionalConditions)
				addCond.PATCH("update", handlers.UpdateAdditionalCondition)
				addCond.DELETE("delete/:id", handlers.DeleteAdditionalCondition)
			}

			accessor := v1.Group("accessor")
			{
				accessor.POST("create", mw.ValidationToken(), handlers.CreateAccessor)
				accessor.GET("read", handlers.ReadAccessors)
				accessor.PATCH("update", handlers.UpdateAccessor)
				accessor.DELETE("delete/:id", handlers.DeleteAccessor)
			}

			munition := v1.Group("munition")
			{
				munition.POST("create", handlers.CreateMunition)
				munition.GET("read", handlers.ReadMunitions)
				munition.PATCH("update", handlers.UpdateMunition)
				munition.DELETE("delete/:id", handlers.DeleteMunition)
			}

			munitionSew := v1.Group("munition_sew")
			{
				munitionSew.POST("create", handlers.CreateMunitionSew)
				munitionSew.GET("read", handlers.ReadMunitionsSew)
				munitionSew.PATCH("update", handlers.UpdateMunitionSew)
				munitionSew.DELETE("delete/:id", handlers.DeleteMunitionSew)
			}

			munitionMod := v1.Group("munition_mod")
			{
				munitionMod.POST("create", handlers.CreateMunitionMod)
				munitionMod.GET("read", handlers.ReadMunitionsMod)
				munitionMod.PATCH("update", handlers.UpdateMunitionMod)
				munitionMod.DELETE("delete/:id", handlers.DeleteMunitionMod)
			}

			pKart := v1.Group("pkart")
			{
				pKart.POST("create", handlers.CreatePKart)
				pKart.GET("read", handlers.ReadPKarts)
			}

			pKartNorm := v1.Group("pkart_norm")
			{
				pKartNorm.POST("create", handlers.CreatePKartNorm)
				pKartNorm.GET("read", handlers.ReadPKartNorms)
			}

			normMunit := v1.Group("norm_munit")
			{
				normMunit.POST("create", handlers.CreateNormMunit)
				normMunit.GET("read", handlers.ReadNormMunit)
				normMunit.PATCH("update")
				normMunit.DELETE("delete/:id")
			}

			normMunitSp := v1.Group("norm_munit_sp")
			{
				normMunitSp.POST("create", handlers.CreateNormMunitSp)
				normMunitSp.GET("read", handlers.ReadNormMunitSp)
			}
		}
	}

	docs.SwaggerInfo.Title = "Автодокументация к приложению Ресурсное обеспечение"
	docs.SwaggerInfo.BasePath = "api/v1"

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Can not start server:", err)
		return err
	}

	return nil
}
