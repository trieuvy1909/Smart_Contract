package routes

import (
	"backend_go/middlewares"
	"backend_go/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	router.Use(cors.New(config))

	userService := services.NewAuthService(db)
	attendanceService := services.NewAttendanceService(db)
	//smartContractService := services.NewSmartContractService(db)

	publicGroup := router.Group("/api")
	privateGroup := router.Group("/api").Use(middlewares.AuthMiddleware())
	{
		publicGroup.POST("/login", userService.Login())
		publicGroup.GET("/logout", userService.Logout())
		publicGroup.POST("/register", userService.Register())
		publicGroup.GET("/time",userService.GetServerTime())
	}
	{
		privateGroup.GET("/get-user-info", userService.GetUserInformation())
		privateGroup.GET("/check-in", attendanceService.CheckIn())
		privateGroup.GET("/check-out", attendanceService.CheckOut())
		privateGroup.POST("/attendance-log-filter", attendanceService.AttendanceLogFilter())
		privateGroup.POST("/attendance-log-detail", attendanceService.AttendanceLogDetail())


		//smart contract
		// privateGroup.GET("/get-value", smartContractService.GetValue())
		// privateGroup.POST("/set-value", smartContractService.SetValue())
	}
	return router
}
