package services

import (
	"backend_go/dtos"
	"backend_go/entities"
	"backend_go/repositories"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db}
}

func (s *AuthService) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		UserRepository := repositories.NewUserRepository(s.db)
		var form dtos.LoginDto
		reps := entities.BaseResponse{}

		// Lấy tham số từ request body
		if err := c.ShouldBind(&form); err != nil {
			reps.Status = 1
			reps.Message = "Lỗi lấy tham số"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		if form.Email == "" {
			reps.Status = 1
			reps.Message = "Thiếu tham số email"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		if form.Password == "" {
			reps.Status = 1
			reps.Message = "Thiếu tham số mật khẩu"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		user, token, message := UserRepository.Login(c, form)

		if message != "" {
			reps.Status = 1
			reps.Message = message
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		reps.Status = 0
		reps.Message = "Thành công"
		reps.Data = gin.H{
			"token": token,
			"infor": user,
		}
		c.JSON(http.StatusOK, reps)
	}
}

func (s *AuthService) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Del("Authorization")
		c.Request.Header.Del("Authorization")
		c.SetCookie("token_"+os.Getenv("IP_EC2"), "", -1, "/", os.Getenv("IP_EC2"), false, true)
		reps := entities.BaseResponse{}
		reps.Status = 0
		reps.Message = "Đăng xuất thành công"
		reps.Data = ""
		c.JSON(http.StatusOK, reps)
	}
}
func (s *AuthService) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		UserRepository := repositories.NewUserRepository(s.db)
		var form dtos.User
		reps := entities.BaseResponse{}

		// Lấy tham số từ request body
		if err := c.ShouldBind(&form); err != nil {
			reps.Status = 1
			reps.Message = "Lỗi lấy tham số"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}

		if form.Email == "" {
			reps.Status = 1
			reps.Message = "Thiếu tham số email"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		if form.Fullname == "" {
			reps.Status = 1
			reps.Message = "Thiếu tham số họ và tên"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		if form.Password == "" {
			reps.Status = 1
			reps.Message = "Thiếu tham số mật khẩu"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		if form.ConfirmPassword == "" {
			reps.Status = 1
			reps.Message = "Thiếu tham số xác nhận mật khẩu"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		if form.Password != form.ConfirmPassword {
			reps.Status = 1
			reps.Message = "Mật khẩu không khớp"
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		user, message := UserRepository.Register(c, form)

		if message != "" {
			reps.Status = 1
			reps.Message = message
			reps.Data = ""
			c.JSON(http.StatusBadRequest, reps)
			return
		}
		reps.Status = 0
		reps.Message = "Đăng kí tài khoản thành công."
		reps.Data = user
		c.JSON(http.StatusOK, reps)
	}
}
func (s *AuthService) GetServerTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		reps := entities.BaseResponse{}
		reps.Status = 0
		reps.Message = "Thành công"
		reps.Data = time.Now().Format("15:04:05 - 02/01/2006")
		c.JSON(http.StatusOK, reps)
	}
}

var jwtKey = os.Getenv("SECRET_KEY")

func (s *AuthService) GetUserInformation() gin.HandlerFunc {
	return func(c *gin.Context) {
		UserRepository := repositories.NewUserRepository(s.db)
		reps := entities.BaseResponse{}
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token, _ = c.Cookie("token_" + os.Getenv("IP_EC2"))

		}
		if token == "" {
			reps.Status = 0
			reps.Message = "Token không tìm thấy"
			reps.Data = ""
		} else {
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
				reps.Status = 0
				reps.Message = "Token lỗi"
				reps.Data = ""
			} else {
				user := UserRepository.GetUserInformation(c, payload["user_id"].(string))
				reps.Status = 0
				reps.Message = "Thông tin tài khoản thành công"
				reps.Data = user
			}
		}
		c.JSON(http.StatusOK, reps)
	}
}
