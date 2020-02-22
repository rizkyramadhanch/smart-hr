package api 

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"smart-hr/modules/components/repositories"
	"smart-hr/modules/components/models"
)

type ComponentController struct{}

func(u *ComponentController) GetByID(ctx *gin.Context){
	repo := repositories.ComponentRepositories{}
	componentID := ctx.Param("id")
	i, e := strconv.Atoi(componentID)
	fmt.Println("received param @ component controller ", i)
	if e != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : e,
		})
		return
	}
	result, err := repo.GetComponent(i)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : err.Error(),
		})
		return
	} else {
		ctx.JSON(200,gin.H{
			"status" : "success",
			"messages" : result,
		})
		return
	}
}

func(u *ComponentController) GetComponentValue(ctx *gin.Context){
	repo := repositories.ComponentRepositories{}
	valID := ctx.Param("id")
	fmt.Println(valID)
	i, e := strconv.Atoi(valID)
	if e != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : e,
		})
		return
	}
	result, err := repo.GetComponentValue(i)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : err.Error(),
		})
		return
	} else {
		ctx.JSON(200,gin.H{
			"status" : "success",
			"messages" : result,
		})
		return
	}
}

func(u *ComponentController) AddComponentValue(ctx *gin.Context){
	repo := repositories.ComponentRepositories{}
	form := models.ValueComponent{}
	errBinding := ctx.ShouldBindJSON(&form)
	if errBinding != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : errBinding,
		})
		return
	}
	
	result, err := repo.AddComponentValue(form)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : err.Error(),
		})
		return
	} else {
		ctx.JSON(200,gin.H{
			"status" : "success",
			"messages" : result,
		})
		return
	}
}

func(u *ComponentController) EditComponentValue(ctx *gin.Context){
	repo := repositories.ComponentRepositories{}
	form := models.Ammount{}
	errBinding := ctx.ShouldBindJSON(&form)
	if errBinding != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : errBinding,
		})
		return
	}
	
	result, err := repo.EditComponentValue(form)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : err.Error(),
		})
		return
	} else {
		ctx.JSON(200,gin.H{
			"status" : "success",
			"messages" : result,
		})
		return
	}
}