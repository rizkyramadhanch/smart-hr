package main

import(
	"github.com/gin-gonic/gin"
	"smart-hr/config"
	"smart-hr/modules/users/api"
	Reports "smart-hr/modules/reports/api"
	Employee "smart-hr/modules/employees/api"
	Company "smart-hr/modules/companies/api"
	"github.com/gin-contrib/cors"
	"fmt"
)

func main(){
	config.Init()
	r := gin.Default()
	
	UserController := api.UserController{}
	ReportController := Reports.ReportsController{}
	EmployeeController := Employee.EmployeeController{}
	CompanyController := Company.CompanyController{}

	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers")
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(cors.New(corsConfig))


	r.POST("/auth", UserController.Login)
	r.POST("/logout/:id", UserController.Logout)
	
	r.POST("/users", UserController.Add)
	r.GET("/users", UserController.GetAll)
	
	r.GET("/employees", EmployeeController.GetAll)
	r.POST("/employees/add", EmployeeController.Add)
	
	r.GET("/reports/bpjs", ReportController.ShowBPJSReport)
	r.GET("/reports/bpjs/:id", ReportController.ShowBPJSReportByID)
	//bpjs/id/month

	//pph
	//pph
	//pph
	
	r.POST("/companies/add", CompanyController.Add)
	r.GET("/companies", CompanyController.GetAll)
	
	fmt.Println("service running on port " + config.Host + ":" + config.Port)
	r.Run(":" + config.Port)
}