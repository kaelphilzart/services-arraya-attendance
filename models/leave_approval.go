package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

type LeaveApprovalModel struct{}

// One ...
func (m LeaveApprovalModel) One(id string) ([]interType.LeaveApproval, error) {
	var leaveApproval []interType.LeaveApproval
	_, err := db.GetDB().Select(&leaveApproval, `
        SELECT 
            la.id,
            la.level,
            CASE 
                WHEN u.id IS NOT NULL THEN jsonb_build_object(
                    'id', u.id, 
                    'Pengaju', jsonb_build_object(
                        'id', u.id,
                        'name', u.name
                    )
                )
                ELSE NULL 
            END AS leave,
			CASE
                WHEN us.id IS NOT NULL THEN jsonb_build_object(
                    'id', us.id,
                    'name', us.name
                )
                ELSE NULL
            END AS approved,
            la.status,
            la.note,
            la.approved_at,
            la.created_at
        FROM sc_attendance.leave_approval la
        LEFT JOIN sc_attendance.leave l ON l.id = la.leave_id 
        LEFT JOIN sc_users.users u ON u.id = l.user_id
        LEFT JOIN sc_users.users us ON us.id = la.approved_by
        WHERE la.leave_id = $1
        ORDER BY la.created_at DESC
    `, id)
	return leaveApproval, err
}

// All ...
func (m LeaveApprovalModel) All() (leaveApproval []interType.LeaveApproval, err error) {
	qs := `
 SELECT 
            la.id,
            la.level,
            CASE 
                WHEN u.id IS NOT NULL THEN jsonb_build_object(
                    'id', u.id, 
                    'Pengaju', jsonb_build_object(
                        'id', u.id,
                        'name', u.name
                    )
                )
                ELSE NULL 
            END AS leave,
			CASE
                WHEN us.id IS NOT NULL THEN jsonb_build_object(
                    'id', us.id,
                    'name', us.name
                )
                ELSE NULL
            END AS approved_by,
            la.status,
            la.note,
            la.approved_at,
            la.created_at
        FROM sc_attendance.leave_approval la
        LEFT JOIN sc_attendance.leave l ON l.id = la.leave_id 
        LEFT JOIN sc_users.users u ON u.id = l.user_id
        LEFT JOIN sc_users.users us ON us.id = la.approved_by
		order by la.created_at desc`
	_, err = db.GetDB().Select(&leaveApproval, qs)
	return leaveApproval, err
}


