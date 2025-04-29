package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type BranchModel struct{}

// One ...
func (m BranchModel) One(id string) (branch interType.Branch, err error) {
	err = db.GetDB().SelectOne(&branch, `
		SELECT 
			b.id,
			b.name,
			b.address,
			b.contact,
			jsonb_build_object(
				'id', c.id,
				'name', c.name
			) AS company
		FROM sc_user.branch b
		LEFT JOIN sc_user.company c 
			ON c.id = b.company_id
		WHERE b.id = $1
		ORDER BY b.created_at DESC
	`, id)
	return branch, err
}


// All ...
func (m BranchModel) All() (branch []interType.Branch, err error) {
	qs := `select b.id, b.name, b.address, b.contact,
	jsonb_build_object('id', c.id, 'name', c.name) AS company
	from sc_user.branch b 
	left join sc_user.company c on c.id = b.company_id 
	order by b.created_at desc`

	_, err = db.GetDB().Select(&branch, qs)
	return branch, err
}
