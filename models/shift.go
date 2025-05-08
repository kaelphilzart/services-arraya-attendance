package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type ShiftModel struct{}

// One ...
func (m ShiftModel) One(id string) (shift interType.Shift, err error) {
	err = db.GetDB().SelectOne(&shift, `
	SELECT 
		s.id,
		s."name",
		CASE 
				WHEN c.id IS NOT NULL THEN jsonb_build_object('id', c.id, 'name', c.name, 'address', c.address, 'contact', c.contact)
				ELSE NULL 
		END AS company,
		s.start_time,
		s.end_time,
		s.created_at, 
		s.updated_at
	FROM sc_attendance.shift s 
	LEFT JOIN sc_users.company c ON c.id = s.company_id
	where s.id = $1
	order by s.created_at desc`, id)
	return shift, err
}

// All ...
func (m ShiftModel) All() (shift []interType.Shift, err error) {
	qs := `
	SELECT 
		s.id,
		s."name",
		CASE 
				WHEN c.id IS NOT NULL THEN jsonb_build_object('id', c.id, 'name', c.name, 'address', c.address, 'contact', c.contact)
				ELSE NULL 
		END AS company,
		s.start_time,
		s.end_time,
		s.created_at, 
		s.updated_at
	FROM sc_attendance.shift s 
	LEFT JOIN sc_users.company c ON c.id = s.company_id
	order by s.created_at desc`
	_, err = db.GetDB().Select(&shift, qs)
	return shift, err
}
