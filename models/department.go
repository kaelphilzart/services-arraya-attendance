package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type DepartmentModel struct{}

// One ...
func (m DepartmentModel) One(id string) (department interType.Department, err error) {
	err = db.GetDB().SelectOne(&department, `
	SELECT 
		d.id,
		d."name",
		CASE 
				WHEN c.id IS NOT NULL THEN jsonb_build_object('id', c.id, 'name', c.name, 'address', c.address, 'contact', c.contact)
				ELSE NULL 
		END AS company,
		CASE 
				WHEN b.id IS NOT NULL THEN jsonb_build_object('id', b.id, 'name', b.name, 'address', b.address, 'contact', b.contact)
				ELSE NULL 
		END AS branch,
		CASE 
        WHEN u.id IS NOT NULL THEN jsonb_build_object(
            'id', u.id, 
            'name', u.name, 
            'email', u.email,
            'photo', up.photo
        )
        ELSE NULL 
   		END AS director,
		d.created_at, 
		d.updated_at
	FROM sc_users.department d 
	LEFT JOIN sc_users.branch b ON b.id = d.branch_id
	LEFT JOIN sc_users.company c ON c.id = d.company_id
	LEFT JOIN sc_users.users u ON u.id = d.director_id
	LEFT JOIN sc_users.user_profile up ON up.user_id = u.id
	where d.id = $1
	order by d.created_at desc`, id)
	return department, err
}

// All ...
func (m DepartmentModel) All() (department []interType.Department, err error) {
	qs := `
		SELECT 
		d.id,
		d."name",
		CASE 
				WHEN c.id IS NOT NULL THEN jsonb_build_object('id', c.id, 'name', c.name, 'address', c.address, 'contact', c.contact)
				ELSE NULL 
		END AS company,
		CASE 
				WHEN b.id IS NOT NULL THEN jsonb_build_object('id', b.id, 'name', b.name, 'address', b.address, 'contact', b.contact)
				ELSE NULL 
		END AS branch,
		CASE 
        WHEN u.id IS NOT NULL THEN jsonb_build_object(
            'id', u.id, 
            'name', u.name, 
            'email', u.email,
            'photo', up.photo
        )
        ELSE NULL 
   		END AS director,
		d.created_at, 
		d.updated_at
	FROM sc_users.department d 
	LEFT JOIN sc_users.branch b ON b.id = d.branch_id
	LEFT JOIN sc_users.company c ON c.id = d.company_id
	LEFT JOIN sc_users.users u ON u.id = d.director_id
	LEFT JOIN sc_users.user_profile up ON up.user_id = u.id
	order by d.created_at desc`
	_, err = db.GetDB().Select(&department, qs)
	return department, err
}
