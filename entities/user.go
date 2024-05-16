package entities

type User struct {
    ID         string `gorm:"column:id;primaryKey;"`
    Email      string `gorm:"column:email"`
    Password   string `gorm:"column:password"`
    Fullname   string `gorm:"column:fullname"`
    Position   string `gorm:"column:position;"`
    Department string `gorm:"column:department;"`
}