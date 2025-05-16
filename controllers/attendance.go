package controllers

import (
	"fmt"
	"log"
	"net/http"
	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"
	"services-arraya-attendance/models"
	"time"

	"services-arraya-attendance/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

type AttendanceController struct {}

var attendanceModel = new(models.AttendanceModel)
var attendanceForm = new(forms.AttendanceForm)

// Attendance In
// func (ctrl AttendanceController) AttendanceIn(c *gin.Context) {
// 	var form forms.AttendanceInForm

// 	// validasi
// 	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
// 		message := attendanceForm.AttendanceIn(validationErr)

// 		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
// 		return
// 	}

// 	form.UserID = getUserID(c)
// 	form.Date = time.Now().Truncate(24 * time.Hour)

// 	// parsing user id to UUID type
// 	parsedUsrId, parseErr := uuid.Parse(getUserID(c))
// 	if parseErr != nil {
// 		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parsing uuid", nil)
// 		return
// 	}
	
// 	// cek attendance today
// 	exists, err := models.AttendanceModel{}.ExistsTodayIn(parsedUsrId, form.Date)
// 	if err != nil {
// 		standarizedResponse(c, true, http.StatusInternalServerError, "Failed to check attendance", nil)
// 		return
// 	}
// 	if exists {
// 		standarizedResponse(c, true, http.StatusConflict, "You have already clocked in today", nil)
// 		return
// 	}

// 	// insert database
// 	id, err := models.FlexibleInsert("sc_attendance.attendance", form, "id")
// 	if err != nil {
// 		fmt.Println(err)
// 		standarizedResponse(c, true, http.StatusNotAcceptable, "Attendance Status could not be created", nil)
// 		return
// 	}

// 	// logging activity
// 	actLog := &interType.LogActivity{
// 		UserID: parsedUsrId, Type: "Attendance", Detail: "Attendance in",
// 	}
// 	go models.LogActivity(actLog)
// 	// ------------------------------

// 	standarizedResponse(c, false, http.StatusOK, "Attendance created", gin.H{"id": id})
// }

func (ctrl AttendanceController) AttendanceIn(c *gin.Context) {
	// Ambil data form biasa
	userID := getUserID(c)
	checkInTime := c.PostForm("chek_in_time")
	latitudeIn := c.PostForm("latitude_in")
	longitudeIn := c.PostForm("longitude_in")

	// Ambil file foto
	file, fileHeader, err := c.Request.FormFile("photo_in")
	if err != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Photo is required", nil)
		return
	}
	defer file.Close()

	// Upload ke Cloudinary
	photoURL, err := utils.UploadToCloudinary(file, fileHeader)
	if err != nil {
		standarizedResponse(c, true, http.StatusInternalServerError, "Failed to upload photo", nil)
		return
	}

	// Build form struct
	form := forms.AttendanceInForm{
		UserID:      userID,
		Date:        time.Now().Truncate(24 * time.Hour),
		ChekInTime:  checkInTime,
		LatitudeIn:  latitudeIn,
		LongitudeIn: longitudeIn,
		PhotoIn:     photoURL, // ini yang disimpan ke DB
	}

	// Validasi (opsional bisa pakai validator manual di sini)
	// Cek apakah sudah absen hari ini
	parsedUsrId, err := uuid.Parse(userID)
	if err != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Invalid UUID", nil)
		return
	}

	exists, err := models.AttendanceModel{}.ExistsTodayIn(parsedUsrId, form.Date)
	if err != nil {
		standarizedResponse(c, true, http.StatusInternalServerError, "Failed to check attendance", nil)
		return
	}
	if exists {
		standarizedResponse(c, true, http.StatusConflict, "You have already clocked in today", nil)
		return
	}

	// Simpan ke DB
	id, err := models.FlexibleInsert("sc_attendance.attendance", form, "id")
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Failed to save attendance", nil)
		return
	}

	// Logging
	actLog := &interType.LogActivity{
		UserID: parsedUsrId, Type: "Attendance", Detail: "Attendance in",
	}
	go models.LogActivity(actLog)

	standarizedResponse(c, false, http.StatusOK, "Attendance created", gin.H{"id": id})
}


// Attendance out 
func (ctrl AttendanceController) AttendanceOut(c *gin.Context) {
	var form forms.AttendanceOutForm
	userId := getUserID(c)

	// validasi
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := attendanceForm.AttendanceOut(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	// cek attendance
	data, err2 := attendanceModel.OneByUserId(userId)
	if err2 != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Attendance not found", nil)
		return
	}

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: data.ID.String(),
	}

	// insert database
	err := models.FlexibleUpdate("sc_attendance.attendance", form, cond, "id")
	if err != nil {
		fmt.Println(err)

		standarizedResponse(c, true, http.StatusNotAcceptable, "Attendance could not be updated", nil)
		return
	}

	// parsing user id to UUID type
	parsedUsrId, parseErr := uuid.Parse(getUserID(c))
	if parseErr != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parsing uuid", nil)
		return
	}

	// logging activity
	actLog := &interType.LogActivity{
		UserID: parsedUsrId, Type: "Attendance", Detail: "Attendance Out",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Attendance Out", nil)
}

// All
func (ctrl AttendanceController) All(c *gin.Context) {
	results, err := attendanceModel.All()
	if err != nil {
		 log.Println("AttendanceModel.All error:", err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get absent", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// One
func (ctrl AttendanceController) One(c *gin.Context) {
	id := c.Param("id")

	data, err := attendanceModel.One(id)
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Absen", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// One By User Id
func (ctrl AttendanceController) OneByUserId(c *gin.Context) {
	id := getUserID(c)

	data, err := attendanceModel.OneByUserId(id)
	if err != nil {
		// Check if the error message contains "no rows in result set"
		if err.Error() == "sql: no rows in result set" {
			standarizedResponse(c, false, http.StatusOK, "No absence record found", map[string]interface{}{})
			return
		}

		// Log other errors and return an internal server error
		fmt.Println("Error retrieving absence data:", err)
		standarizedResponse(c, true, http.StatusInternalServerError, "Could not get Absen", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// History
func (ctrl AttendanceController) History(c *gin.Context) {
	id := getUserID(c)

	data, err := attendanceModel.History(id)
	if err != nil {
		// Check if the error message contains "no rows in result set"
		if err.Error() == "sql: no rows in result set" {
			standarizedResponse(c, false, http.StatusOK, "No absence record found", map[string]interface{}{})
			return
		}

		// Log other errors and return an internal server error
		fmt.Println("Error retrieving absence data:", err)
		standarizedResponse(c, true, http.StatusInternalServerError, "Could not get Absen", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}
