package api

import (
	"github.com/gin-gonic/gin"
	"smart-hr/modules/reports/repositories"
	"strconv"
)

type ReportsController struct{}

func (u *ReportsController) ShowBPJSReport(ctx *gin.Context) {
	repo := repositories.ReportRepositories{}
	result, err := repo.ShowBPJSReport()
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"messages": err.Error(),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"status":   "success",
			"messages": result,
		})
		return
	}
}

func (u *ReportsController) ShowBPJSReportByID(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	repo := repositories.ReportRepositories{}
	result, err := repo.ShowBPJSReportByID(userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"messages": err.Error(),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"status":   "success",
			"messages": result,
		})
		return
	}
}

//start here..
func (u *ReportsController) SallaryReport(ctx *gin.Context) {
	repo := repositories.ReportRepositories{}
	result, err := repo.SallaryReportAll()
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"messages": err.Error(),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"status":   "success",
			"messages": result,
		})
		return
	}
}

func (u *ReportsController) SallaryReportByID(ctx *gin.Context) {
	repo := repositories.ReportRepositories{}
	employeeID := ctx.Param("id")
	i, e := strconv.Atoi(employeeID)
	if e != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"messages": e.Error(),
		})
		return
	}
	result, err := repo.SallaryReport(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"messages": err.Error(),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"status":   "success",
			"messages": result,
		})
		return
	}
}
