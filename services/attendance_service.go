package services

import (
	"backend_go/dtos"
	"backend_go/entities"
	"backend_go/repositories"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type AttendanceService struct {
	db *gorm.DB
}

func NewAttendanceService(db *gorm.DB) *AttendanceService {
	return &AttendanceService{db}
}

func (s *AttendanceService) CheckIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		AttendanceRepository := repositories.NewAttendanceRepository(s.db)
		var form dtos.AttendanceDto
		reps := entities.BaseResponse{}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token, _ = c.Cookie("token_" + os.Getenv("IP_EC2"))

		}
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}
		tokenObject, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		var payload jwt.MapClaims
		if claims, ok := tokenObject.Claims.(jwt.MapClaims); ok && tokenObject.Valid {
			payload = claims
		}
		if payload == nil {
			reps.Status = 1
			reps.Message = "Token lỗi"
			reps.Data = ""
			return
		} else {
			form.UserID = payload["user_id"].(string)
			attendance, message := AttendanceRepository.CheckIn(c, form)
			if message != "" {
				reps.Status = 1
				reps.Message = message
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			reps.Status = 0
			reps.Message = "Thành công"
			reps.Data = attendance
			c.JSON(http.StatusOK, reps)
			return
		}
	}
}
func (s *AttendanceService) CheckOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		AttendanceRepository := repositories.NewAttendanceRepository(s.db)
		var form dtos.AttendanceDto
		reps := entities.BaseResponse{}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token, _ = c.Cookie("token_" + os.Getenv("IP_EC2"))

		}
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}
		tokenObject, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		var payload jwt.MapClaims
		if claims, ok := tokenObject.Claims.(jwt.MapClaims); ok && tokenObject.Valid {
			payload = claims
		}
		if payload == nil {
			reps.Status = 1
			reps.Message = "Token lỗi"
			reps.Data = ""
			return
		} else {
			form.UserID = payload["user_id"].(string)
			attendance, message := AttendanceRepository.CheckOut(c, form)
			if message != "" {
				reps.Status = 1
				reps.Message = message
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			reps.Status = 0
			reps.Message = "Thành công"
			reps.Data = attendance
			c.JSON(http.StatusOK, reps)
			return
		}
	}
}
func (s *AttendanceService) AttendanceLogFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		AttendanceRepository := repositories.NewAttendanceRepository(s.db)
		var form dtos.AttendanceLogDto
		reps := entities.BaseResponse{}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token, _ = c.Cookie("token_" + os.Getenv("IP_EC2"))

		}
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}
		tokenObject, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		var payload jwt.MapClaims
		if claims, ok := tokenObject.Claims.(jwt.MapClaims); ok && tokenObject.Valid {
			payload = claims
		}
		if payload == nil {
			reps.Status = 1
			reps.Message = "Token lỗi"
			reps.Data = ""
			return
		} else {
			// Lấy tham số từ request body
			if err := c.ShouldBind(&form); err != nil {
				reps.Status = 1
				reps.Message = "Lỗi lấy tham số"
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			if form.StartTime != "" && form.EndTime == ""{
				reps.Status = 1
				reps.Message = "Thiếu tham số end_time"
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			form.UserID = payload["user_id"].(string)
			attendance, message := AttendanceRepository.AttendanceLogFilter(c, form)
			if message != "" {
				reps.Status = 1
				reps.Message = message
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			reps.Status = 0
			reps.Message = "Thành công"
			reps.Data = attendance
			c.JSON(http.StatusOK, reps)
			return
		}
	}
}
func (s *AttendanceService) AttendanceLogDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		AttendanceRepository := repositories.NewAttendanceRepository(s.db)
		var form dtos.AttendanceLogDto
		reps := entities.BaseResponse{}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token, _ = c.Cookie("token_" + os.Getenv("IP_EC2"))

		}
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}
		tokenObject, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		var payload jwt.MapClaims
		if claims, ok := tokenObject.Claims.(jwt.MapClaims); ok && tokenObject.Valid {
			payload = claims
		}
		if payload == nil {
			reps.Status = 1
			reps.Message = "Token lỗi"
			reps.Data = ""
			return
		} else {
			// Lấy tham số từ request body
			if err := c.ShouldBind(&form); err != nil {
				reps.Status = 1
				reps.Message = "Lỗi lấy tham số"
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			if form.AttendanceID == 0{
				reps.Status = 1
				reps.Message = "Thiếu tham số attendance_id"
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			form.UserID = payload["user_id"].(string)
			attendance, message := AttendanceRepository.AttendanceLogDetail(c, form)
			if message != "" {
				reps.Status = 1
				reps.Message = message
				reps.Data = ""
				c.JSON(http.StatusBadRequest, reps)
				return
			}
			reps.Status = 0
			reps.Message = "Thành công"
			reps.Data = attendance
			c.JSON(http.StatusOK, reps)
			return
		}
	}
}