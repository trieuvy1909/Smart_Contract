package entities

type Attendance struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	UserID    string  `gorm:"column:user_id;"`
	CheckIn   string  `gorm:"column:check_in"`
	CheckOut  string  `gorm:"column:check_out" `
	Location  string  `gorm:"column:location"`
	OverTime  float64 `gorm:"column:over_time"`
	Bonus     float64 `gorm:"column:bonus"`
	Salary    float64 `gorm:"column:salary"`
	TotalHour float64 `gorm:column:total_hour"`
}
type AttendanceFilter struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	UserID    string  `gorm:"column:user_id;"`
	CheckIn   string  `gorm:"column:check_in"`
	CheckOut  string  `gorm:"column:check_out" `
	TotalHour float64 `gorm:column:total_hour"`
}
type AttendanceDetail struct {
	ID         int     `gorm:"primaryKey;autoIncrement"`
	UserID     string  `gorm:"column:user_id"`
	Fullname   string  `gorm:"column:fullname"`
	Position   string  `gorm:"column:position"`
	Department string  `gorm:"column:department"`
	CheckIn    string  `gorm:"column:check_in"`
	CheckOut   string  `gorm:"column:check_out" `
	Location   string  `gorm:"column:location"`
	OverTime   float64 `gorm:"column:over_time"`
	Bonus      float64 `gorm:"column:bonus"`
	Salary     float64 `gorm:"column:salary"`
	TotalHour  float64 `gorm:column:total_hour"`
}
