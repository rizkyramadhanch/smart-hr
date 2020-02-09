package api 

import (
	"github.com/gin-gonic/gin"
	_ "strconv"
	"smart-hr/modules/companies/repositories"
	"smart-hr/modules/companies/models"
)

type CompanyController struct{}

func(u *CompanyController) Add(ctx *gin.Context){
	repo := repositories.CompanyRepositories{}
	form := models.Company{}
	
	err := ctx.ShouldBindJSON(&form)
	if(err != nil) {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : err.Error(),
		})
		return
	}
	
	result, err := repo.AddCompany(form)
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

func(u *CompanyController) GetAll(ctx *gin.Context){
	repo := repositories.CompanyRepositories{}
	
	result, err := repo.GetAll()
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