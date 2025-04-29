package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type TypeLeaveModel struct{}

// One ...
func (m TypeLeaveModel) One(id string) (typeLeave interType.TypeLeave, err error) {
	err = db.GetDB().SelectOne(&typeLeave, `
	select 
	tl.id,
	tl.code,
	tl.name, 
	tl.created_at,
	tl.updated_at 
	from sc_attendance.type_leave tl
	where tl.id = $1
	order by tl.created_at desc`, id)  
	return typeLeave, err
}

// All ...
func (m TypeLeaveModel) All() (typeLeave []interType.TypeLeave, err error) {
	qs := `select 
	tl.id,
	tl.code,
	tl.name, 
	tl.created_at,
	tl.updated_at 
	from sc_attendance.type_leave tl
	order by tl.created_at desc`
	_, err = db.GetDB().Select(&typeLeave, qs)
	return typeLeave, err
}
