package dtos

type LoginDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
type User struct {
    ID         string `form:"id"`
    Email      string `form:"email"`
    Password   string `form:"password"`
    Fullname   string `form:"fullname"`
    Position   string `form:"position"`
    Department string `form:"department"`
	ConfirmPassword string `form:"confirm_password"`
}