package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type CompanyModel struct{}

// One ...
func (m CompanyModel) One(id string) (Company interType.Company, err error) {
	err = db.GetDB().SelectOne(&Company, "select * from sc_user.company c where c.id = $1 order by c.created_at desc", id)
	return Company, err
}

// All ...
func (m CompanyModel) All() (Company []interType.Company, err error) {
	qs := `select * from sc_user.company c order by c.created_at desc`

	_, err = db.GetDB().Select(&Company, qs)
	return Company, err
}
