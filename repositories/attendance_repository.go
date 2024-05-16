package repositories

import (
	"backend_go/dtos"
	"backend_go/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{db}
}
func (r *AttendanceRepository) CheckIn(c *gin.Context, form dtos.AttendanceDto) (*entities.Attendance, string) {

	attendance := &entities.Attendance{}
	attendance.UserID = form.UserID;
	attendance.CheckIn = time.Now().Format("2006/01/02 - 15:04:05")
	if form.Location == ""{
		attendance.Location = "Vinaconex Building - 47 Dien Bien Phu, Da kao, Dictrict 1, Ho Chi Minh city"
	}
	result := r.db.Create(&attendance)
	if result.Error != nil {
		return attendance, "Lỗi thêm dữ liệu vào database"
	}

	return attendance, ""
}
func (r *AttendanceRepository) CheckOut(c *gin.Context, form dtos.AttendanceDto) (*entities.Attendance, string) {

	var attendance *entities.Attendance
	resultCheckExist := r.db.Where("user_id = ?", form.UserID).Order("id DESC").First(&attendance)
	if resultCheckExist.Error != nil {
		return nil, "Không tồn tại dữ liệu check in tương ứng"
	}
	attendance.CheckOut  = time.Now().Format("2006/01/02 - 15:04:05")
	checkIn, _ := time.Parse("2006/01/02 - 15:04:05",attendance.CheckIn)
	checkOut, _ := time.Parse("2006/01/02 - 15:04:05",attendance.CheckOut)
	attendance.TotalHour = checkOut.Sub(checkIn).Hours()
	result := r.db.Model(&entities.Attendance{}).Where("id = ?", attendance.ID).Updates(&attendance)
	if result.Error != nil {
		return attendance, "Lỗi cập nhật dữ liệu trong cơ sở dữ liệu"
	}

	return attendance, ""
}
func (r *AttendanceRepository) AttendanceLogFilter(c *gin.Context, form dtos.AttendanceLogDto) ([]entities.AttendanceFilter, string) {

	var attendanceLog []entities.AttendanceFilter
	form.UserID = "%"+form.UserID+"%"
	if form.StartTime == "" {
		form.StartTime = "0000/00/00 - 00:00:00"
	}
	if form.EndTime == "" {
		form.EndTime = "9999/99/99 - 99:99:99"
	}
	query:="select at.id,at.user_id,us.fullname,us.department,us.position,at.check_in,at.check_out, at.total_hour from attendances at left join users us on at.user_id = us.id where at.user_id like ? and at.check_in >= ? and at.check_out <= ?"
	if err := r.db.Raw(query, form.UserID,form.StartTime,form.EndTime).Scan(&attendanceLog).Error; err != nil {
		return attendanceLog, "Không truy vấn dữ liệu điểm danh"
	}
	return attendanceLog, ""
}
func (r *AttendanceRepository) AttendanceLogDetail(c *gin.Context, form dtos.AttendanceLogDto) (*entities.AttendanceDetail, string) {

	var attendanceLogDetail *entities.AttendanceDetail
	query:="select at.*,us.fullname,us.department,us.position from attendances at left join users us on at.user_id = us.id where at.id = ?"
	if err := r.db.Raw(query, form.AttendanceID).Scan(&attendanceLogDetail).Error; err != nil {
		return attendanceLogDetail, "Không truy vấn dữ liệu điểm danh"
	}
	return attendanceLogDetail, ""
}