package dtos

type AttendanceDto struct {
	ID        int     `form:"id"`
	UserID    string  `form:"user_id" json:"user_id"`
	CheckIn   string  `form:"check_in"`
	CheckOut  string  `form:"check_out"`
	Location  string  `form:"location"`
	OverTime  float64 `form:"over_time"`
	Bonus     float64 `form:"bonus"`
	Salary    float64 `form:"salary"`
	TotalHour float64 `form:total_hour"`
}

type AttendanceLogDto struct {
	AttendanceID int    `form:"attendance_id" json:"attendance_id"`
	UserID       string `form:"user_id"`
	StartTime    string `form:"start_time"`
	EndTime      string `form:"end_time"`
}