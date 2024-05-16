package middlewares

import (
	"backend_go/entities"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
 
 
 var jwtKey = os.Getenv("SECRET_KEY")
 
 func CreateToken(user *entities.User, c *gin.Context) (string, error){
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["fullname"] = user.Fullname
	claims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix()
 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err== nil {
		c.Header("Authorization", tokenString)
		c.SetCookie("token_"+os.Getenv("IP_EC2"), tokenString, 7*24*60*60, "/", os.Getenv("IP_EC2"), false, true)  
	}
	return tokenString,err
}
func getTokenFromIP(c *gin.Context,ip string) string {
	token, _ := c.Cookie("token_"+ ip)  
	return token 
}
 func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	   	token := c.Request.Header.Get("Authorization")
	   	if token == "" {
			ip := os.Getenv("IP_EC2")  
			token := getTokenFromIP(c,ip)  
			_,err := ValidateToken(token)
			if err != "" {
				c.JSON(http.StatusBadRequest, gin.H {
					"status":1,
					"message":err,
					"data":"",
				})
				c.Abort()
				return
			}
	   	}
	   c.Next()
	}
 }

func ValidateToken(tokenString string) (*jwt.Token, string) {
	if tokenString == "" {
		return nil, "Bạn chưa đăng nhập, vui lòng đăng nhập lại."
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Cách tạo token sai, vui lòng đăng nhập lại.")  
		}
		return []byte(jwtKey), nil
	})
	
	if err != nil {
		return nil, "Token không lợp lệ, vui lòng đăng nhập lại."
	}
	
	// Kiểm tra token hợp lệ
	if !token.Valid { 
		return nil, "Token không lợp lệ, vui lòng đăng nhập lại."
	}
	
	now := time.Now().Unix()
    
    if token.Claims.(jwt.MapClaims)["exp"].(float64) < float64(now){
       return nil,"Phiên đăng nhập hết hạn, vui lòng đăng nhập lại."
    }
	return token, ""
}
