package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
	"time"
	uuid "github.com/google/uuid"
)

type AttendanceModel struct{}

// One By User ID...
func (m AttendanceModel) OneByUserId(id string) (attendance interType.Attendance, err error) {
	dateNow := time.Now().Truncate(24 * time.Hour)

	err = db.GetDB().SelectOne(&attendance, `
        select 
            a.id,
            CASE 
                WHEN u.id IS NOT NULL THEN jsonb_build_object('id', u.id, 'name', u.name)
                ELSE NULL 
            END AS user,
            a."date",
            a.chek_in_time,
            a.chek_out_time,
            a.latitude_in,
            a.longitude_in,
            a.latitude_out,
            a.longitude_out,
            a.photo_in,
            a.photo_out,
            a.note,
            a.created_at,
            a.updated_at
        from sc_attendance.attendance a 
        left join sc_users.users u on u.id = a.user_id
        WHERE a.user_id = $1
        AND a.date = $2
        order by a.created_at desc
        LIMIT 1`, id, dateNow)

	return attendance, err
}

// One ...
func (m AttendanceModel) One(id string) ([]interType.Attendance, error) {
	var attendance []interType.Attendance
	_, err := db.GetDB().Select(&attendance, `
        SELECT 
            a.id,
            CASE 
                WHEN u.id IS NOT NULL THEN jsonb_build_object(
                    'id', u.id, 
                    'name', u.name,
                    'profile', jsonb_build_object(
                        'full_name', up.full_name,
                        'photo', up.photo, 
                        'phone', up.phone_number,
                        'birth_date', up.birth_date,
                        'birth_place', up.birth_place,
                        'address', up.address,
                    ),
                    'position', jsonb_build_object(
                        'id', p.id,
                        'name', p.name
                    ),
                      'department', jsonb_build_object(
                        'id', d.id,
                        'name', d.name
                    ),
                    'company', jsonb_build_object(
                        'id', c.id,
                        'name', c.name
                    ),
                    'shift', jsonb_build_object(
                        'id', s.id,
                        'name', s.name,
                        'start_time', s.start_time,
                        'end_time', s.end_time
                    )
                )
                ELSE NULL 
            END AS user,
            a."date",
            a.chek_in_time,
            a.chek_out_time,
            a.latitude_in,
            a.longitude_in,
            a.latitude_out,
            a.longitude_out,
            a.photo_in,
            a.photo_out,
            a.note,
            a.created_at,
            a.updated_at
        FROM sc_attendance.attendance a 
        LEFT JOIN sc_users.users u ON u.id = a.user_id
        LEFT JOIN sc_users.user_profile up ON up.user_id = u.id
        LEFT JOIN sc_users.position p ON p.id = u.position_id
        LEFT JOIN sc_users.department d ON d.id = u.department_id
        LEFT JOIN sc_users.company c ON c.id = d.company_id
        LEFT JOIN sc_attendance.shift s ON s.company_id = c.id
        WHERE a.user_id = $1
        ORDER BY a.created_at DESC
    `, id)
	return attendance, err
}

// All ...
func (m AttendanceModel) All() (attendance []interType.Attendance, err error) {
	qs := `
			SELECT 
            a.id,
            CASE 
                WHEN u.id IS NOT NULL THEN jsonb_build_object(
                    'id', u.id, 
                    'name', u.name,
                    'position', jsonb_build_object(
                        'id', p.id,
                        'name', p.name
                    )
                )
                ELSE NULL 
            END AS user,
            a."date",
            a.chek_in_time,
            a.chek_out_time,
            a.latitude_in,
            a.longitude_in,
            a.latitude_out,
            a.longitude_out,
            a.photo_in,
            a.photo_out,
            a.note,
            a.created_at,
            a.updated_at
        FROM sc_attendance.attendance a 
        LEFT JOIN sc_users.users u ON u.id = a.user_id
        LEFT JOIN sc_users.position p ON p.id = u.position_id
        LEFT JOIN sc_attendance.user_shift us ON us.user_id = u.id
        LEFT JOIN sc_attendance.shift s ON s.id = us.shift_id
		order by a.created_at desc`

	_, err = db.GetDB().Select(&attendance, qs)
	return attendance, err
}

// cek today in
func (m AttendanceModel) ExistsTodayIn(id uuid.UUID, date time.Time) (bool, error) {
	var count int
	err := db.GetDB().SelectOne(&count, `
		SELECT COUNT(1)
		FROM sc_attendance.attendance
		WHERE user_id = $1
		AND date = $2
		AND chek_in_time IS NOT NULL
	`, id, date)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Cek Today Out
func (m AttendanceModel) ExistsTodayOut(id uuid.UUID, date time.Time) (bool, error) {
	var count int
	err := db.GetDB().SelectOne(&count, `
		SELECT COUNT(1)
		FROM sc_attendance.attendance
		WHERE user_id = $1
		AND date = $2
		AND chek_in_time IS NOT NULL
		AND chek_out_time IS NOT NULL
	`, id, date)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// History
func (m AttendanceModel) History(id string) (attendance []interType.Attendance, err error) {
	qs := `
		SELECT 
            a.id,
            CASE 
                WHEN u.id IS NOT NULL THEN jsonb_build_object(
                    'id', u.id, 
                    'name', u.name,
                    'position', jsonb_build_object(
                        'id', p.id,
                        'name', p.name
                    )
                )
                ELSE NULL 
            END AS user,
            a."date",
            a.chek_in_time,
            a.chek_out_time,
            a.latitude_in,
            a.longitude_in,
            a.latitude_out,
            a.longitude_out,
            a.photo_in,
            a.photo_out,
            a.note,
            a.created_at,
            a.updated_at
        FROM sc_attendance.attendance a 
        LEFT JOIN sc_users.users u ON u.id = a.user_id
        LEFT JOIN sc_users.position p ON p.id = u.position_id
        LEFT JOIN sc_attendance.user_shift us ON us.user_id = u.id
        LEFT JOIN sc_attendance.shift s ON s.id = us.shift_id
        WHERE a.user_id = $1
		ORDER BY a.created_at DESC
	`

	_, err = db.GetDB().Select(&attendance, qs, id)
	return attendance, err
}

