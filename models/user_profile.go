package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"

)

// Model ...
type UserProfileModel struct{}

// One ...
func (m UserProfileModel) One(id string) (profileUser interType.UserProfile, err error) {
	err = db.GetDB().SelectOne(&profileUser, `
	SELECT 
		uf.id, 
		uf.full_name,  
		uf.birth_date, 
		uf.birth_place,  
		uf.address, 
		uf.phone_number,
		uf.gender,
		uf.photo,
		CASE 
			WHEN u.id IS NOT NULL THEN jsonb_build_object(
				'user_id', u.id,
				'name', u.name,
				'email', u.email,
				'role', jsonb_build_object(
					'id', r.id,
					'name', r.name
				),
				'company', jsonb_build_object(
					'id', c.id,
					'name', c.name
				),
				'branch', jsonb_build_object(
					'id', b.id,
					'name', b.name
				),
				'position', jsonb_build_object(
					'id', p.id,
					'name', p.name
				),
				'user_shift', jsonb_build_object(
					'id', us.id,
					'shift_name', us.shift_name,
					'start_time', us.start_time,
					'end_time', us.end_time
				)
			)
			ELSE NULL
		END AS user,
		uf.created_at, 
		uf.updated_at,
		uf.deleted_at 
	FROM sc_user.user_profile uf
	LEFT JOIN sc_user.users u ON u.id = uf.user_id
	LEFT JOIN sc_user.user_shift us ON us.user_id = u.id
	LEFT JOIN sc_user.company c ON c.id = u.company_id
	LEFT JOIN sc_user.branch b ON b.id = u.branch_id
	LEFT JOIN sc_user.position p ON p.id = u.position_id
	LEFT JOIN sc_user.role r ON r.id = u.role_id
	WHERE uf.user_id = $1
	LIMIT 1
	`, id)

	return profileUser, err
}
