package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

type LeaveModel struct{}

// One ByUserId...
func (m LeaveModel) One(id string) ([]interType.Leave, error) {
	var leave []interType.Leave
	_, err := db.GetDB().Select(&leave, `
        SELECT 
            l.id,
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
			CASE
                WHEN tl.id IS NOT NULL THEN jsonb_build_object(
                    'id', tl.id,
                    'name', tl.name
                )
                ELSE NULL
            END AS type_leave,
            l.url_photo,
            l.start_date,
            l.end_date,
            l.status,
            l.current_approval_level,
            l.description,
            l.created_at,
            l.updated_at
        FROM sc_attendance.leave l 
        LEFT JOIN sc_users.users u ON u.id = l.user_id
        LEFT JOIN sc_users.position p ON p.id = u.position_id
		LEFT JOIN sc_attendance.type_leave tl ON tl.id = l.type_leave_id
        WHERE l.user_id = $1
        ORDER BY l.created_at DESC
    `, id)
	return leave, err
}

// OneByDepartment ...
func (m LeaveModel) OneByDepartment(departmentId string) ([]interType.Leave, error) {
	var leave []interType.Leave
	_, err := db.GetDB().Select(&leave, `
        SELECT 
            l.id,
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
            CASE
                WHEN tl.id IS NOT NULL THEN jsonb_build_object(
                    'id', tl.id,
                    'name', tl.name
                )
                ELSE NULL
            END AS type_leave,
            l.url_photo,
            l.start_date,
            l.end_date,
            l.status,
            l.current_approval_level,
            l.description,
            l.created_at,
            l.updated_at
        FROM sc_attendance.leave l 
        LEFT JOIN sc_users.users u ON u.id = l.user_id
        LEFT JOIN sc_users.position p ON p.id = u.position_id
        LEFT JOIN sc_attendance.type_leave tl ON tl.id = l.type_leave_id
        WHERE p.department_id = $1
        ORDER BY l.created_at DESC
    `, departmentId)
	return leave, err
}

// All ...
func (m LeaveModel) All() (leave []interType.Leave, err error) {
	qs := `
		SELECT 
            l.id,
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
			CASE
                WHEN tl.id IS NOT NULL THEN jsonb_build_object(
                    'id', tl.id,
                    'name', tl.name
                )
                ELSE NULL
            END AS type_leave,
            l.url_photo,
            l.start_date,
            l.end_date,
            COALESCE(l.status, false) AS status,
            l.current_approval_level,
            l.description,
            l.created_at,
            l.updated_at
        FROM sc_attendance.leave l 
        LEFT JOIN sc_users.users u ON u.id = l.user_id
        LEFT JOIN sc_users.position p ON p.id = u.position_id
		LEFT JOIN sc_attendance.type_leave tl ON tl.id = l.type_leave_id
        WHERE COALESCE(l.status, false) = false
		order by l.created_at desc`
	_, err = db.GetDB().Select(&leave, qs)
	return leave, err
}

// get current level
func (m LeaveModel) GetStatus(LeaveId string) (Leave interType.Leave, err error) {
	err = db.GetDB().SelectOne(&Leave, "SELECT l.id, l.status, l.current_approval_level FROM sc_attendance.leave l WHERE l.leave_id=$1 LIMIT 1", LeaveId)
	return Leave, err
}