package models

import (
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"
)

// Model ...
type LogActivityModel struct{}

// All ...
func (m LogActivityModel) All(params *interType.ParamsLogActivityAll) (logActivity []interType.LogActivity, err error) {
	_, err = db.GetDB().Select(&logActivity, "select la.id, la.user_id, u.'name', la.type, la.detail, u.url_photo from sc_users.log_activity la left join sc_users.users u on u.id = la.user_id where la.user_id = $1 order by la.created_at desc", params.UserID)
	return logActivity, err
}
