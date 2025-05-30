package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type PositionModel struct{}

// One ...
func (m PositionModel) One(id string) (position interType.Position, err error) {
	err = db.GetDB().SelectOne(&position, `
	select 
	p.id,
	p."name",
	p.level, 
	jsonb_build_object('id', d.id, 'name', d.name, 'director', u.name) as department,
	p.created_at,
	p.updated_at 
	from sc_users.position p
	left join sc_users.department d on d.id = p.department_id 
	left join sc_users.users u on u.id = d.director_id
	where p.id = $1
	order by p.created_at desc`, id)  
	return position, err
}

// All ...
func (m PositionModel) All() (department []interType.Position, err error) {
	qs := `select 
	p.id,
	p.name,
	p.level,
	jsonb_build_object(
		'id', d.id,
		'name', d.name,
		'director', u.name,
		'company', jsonb_build_object(
		'id', c.id,
		'name', c.name
		)
	) AS department,
	p.created_at,
	p.updated_at
	from sc_users.position p
	left join sc_users.department d on d.id = p.department_id 
	left join sc_users.company c on c.id = d.company_id
	left join sc_users.users u on u.id = d.director_id
	order by p.created_at desc`

	_, err = db.GetDB().Select(&department, qs)
	return department, err
}
