package repositories

import (
	"backend_go/dtos"
	"backend_go/entities"
	"backend_go/middlewares"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"fmt"
	"strings"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
func generateGUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	guid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	guid = strings.ReplaceAll(guid, "-", "")
	return guid
}
func (r *UserRepository) Login(c *gin.Context, form dtos.LoginDto) (*entities.User, string, string) {

	var token = ""
	var user *entities.User

	err := r.db.Table("users").Where("email = ?", form.Email).First(&user).Error
	if err != nil {
		return user, token, "Email không tồn tại"
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		return user, token, "Mật khẩu sai"
	}
	user.Password = ""
	token, err = middlewares.CreateToken(user, c)
	if token == "" || err != nil {
		return user, token, "Lỗi tạo token"
	}
	return user, token, ""
}
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func (r *UserRepository) Register(c *gin.Context, form dtos.User) (*entities.User, string) {

	var user *entities.User

	err := r.db.Table("users").Where("email = ?", form.Email).First(&user).Error
	if err == nil && user.ID != "" {
		return user, "Email này đã đăng kí tài khoản"
	}

	user.ID = generateGUID()
	user.Email = form.Email
	user.Fullname = form.Fullname
	user.Position = form.Position
	user.Department = form.Department
	user.Password = HashPassword(form.Password)
	
	result := r.db.Create(&user)
	if result.Error != nil {
		return user, "Lỗi thêm dữ liệu vào database"
	}
	
	user.Password = ""
	return user, ""
}
func (r *UserRepository) GetUserInformation(c *gin.Context, user_id string) *entities.User {

	var user *entities.User
	query := "SELECT A.* FROM users A  WHERE A.id = ? LIMIT 1"
	result := r.db.Raw(query, user_id).Scan(&user)
	if result.Error != nil {
		return user
	}
	user.Password = ""

	return user
}
