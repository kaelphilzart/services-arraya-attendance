package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type RoleModel struct{}

// One ...
func (m RoleModel) One(id string) (role interType.Role, err error) {
	err = db.GetDB().SelectOne(&role, `
	SELECT 
		r.id,
		r."name",
		r.slug_name,
		r.created_at, 
		r.updated_at
	FROM sc_users.role r	
	where r.id = $1
	order by r.created_at desc`, id)
	return role, err
}

// All ...
func (m RoleModel) All() (role []interType.Role, err error) {
	qs := `SELECT 
		r.id,
		r."name",
		r.slug_name,
		r.created_at, 
		r.updated_at
	FROM sc_users.role r 
	order by r.created_at desc`
	_, err = db.GetDB().Select(&role, qs)
	return role, err
}
