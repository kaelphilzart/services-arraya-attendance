package models

import (
	"encoding/json"

	"services-arraya-attendance/db"
	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"

	"golang.org/x/crypto/bcrypt"
)

// Model ...
type UserModel struct{}

var authModel = new(AuthModel)

// Login ...
func (m UserModel) Login(form forms.LoginForm) (user interType.User, token interType.Token, err error) {
	qs := `SELECT 
		u.id, 
		u.email, 
		u.password,
		CASE 
			WHEN r.id IS NOT NULL THEN jsonb_build_object('id', r.id, 'name', r.name, 'slug_name', r.slug_name)
			ELSE NULL 
		END AS role
	FROM sc_users.users u
	LEFT JOIN sc_users.role r ON r.id = u.role_id
	WHERE u.email = LOWER($1) AND u.active IS TRUE
	LIMIT 1`
	err = db.GetDB().SelectOne(&user, qs, form.Email)

	if err != nil {
		return user, token, err
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}

	var decodedRole interType.Role
	if user.Role != nil {
		err = json.Unmarshal(*user.Role, &decodedRole)
		if err != nil {
			return user, token, err
		}
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID, decodedRole.SlugName)
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

// ChangePassword ...
// func (m UserModel) ChangePassword(idUsr string, form forms.ChangePasswordForm) (user interType.User, passAct interType.ChangePasswordAct, err error) {

// 	err = db.GetDB().SelectOne(&user, "select tu.id, tu.'password' from sc_users.tb_user tu where tu.id = $1 and tu.active is true limit 1", idUsr)

// 	if err != nil {
// 		return user, passAct, err
// 	}

// 	//Compare the password form and database if match
// 	bytePassword := []byte(form.OldPassword)
// 	byteHashedPassword := []byte(user.Password)

// 	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
// 	if err != nil {
// 		return user, passAct, err
// 	}

// 	//compare the new password and confirm password
// 	if form.ConfirmPassword != form.NewPassword {
// 		return user, passAct, errors.New("something went wrong, please try again later")
// 	}

// 	//create hashed new password
// 	bytePassword = []byte(form.NewPassword)
// 	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
// 	if err != nil {
// 		return user, passAct, errors.New("something went wrong, please try again later")
// 	}

// 	passAct.Password = string(hashedPassword)

// 	return user, passAct, nil
// }

// One By Email ...
func (m UserModel) OneByEmail(email string) (user interType.User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT u.id, u.email, u.name FROM sc_users.users tu WHERE u.email=$1 LIMIT 1", email)
	return user, err
}

// One ...
func (m UserModel) One(id string) (user interType.User, err error) {
	err = db.GetDB().SelectOne(&user, `SELECT 
		u.id, 
		u.name, 
		u.email, 
		tu.password,
		CASE 
			WHEN r.id IS NOT NULL THEN jsonb_build_object('id', r.id, 'name', r.name) 
			ELSE NULL 
		END AS role,
		CASE 
			WHEN c.id IS NOT NULL THEN jsonb_build_object('id', c.id, 'name', c.name) 
			ELSE NULL 
		END AS company,
		CASE 
			WHEN b.id IS NOT NULL THEN jsonb_build_object('id', b.id, 'name', b.name, 'address', b.address, 'contact', b.contact) 
			ELSE NULL 
		END AS branch, 
		CASE 
			WHEN d.id IS NOT NULL THEN jsonb_build_object('id', d.id, 'name', d.name) 
			ELSE NULL 
		END AS department,
		CASE 
			WHEN p.id IS NOT NULL THEN jsonb_build_object('id', p.id, 'name', p.name) 
			ELSE NULL 
		END AS position,
		CASE
			WHEN up.id IS NOT NULL THEN jsonb_build_object('full_name', up.full_name, 'birth_date', up.birth_date, 'birth_place', up.birth_place, 'address', up.address, 'phone_number', up.phone_number, 'gender', up.gender, 'photo', up.photo,)
			ELSE NULL
		END AS profile,
		u.active,
		u.created_at, 
		u.updated_at,
		u.deleted_at 
	FROM sc_users.users u
	LEFT JOIN sc_users.position p ON p.id = u.position_id
	LEFT JOIN sc_users.branch b ON b.id = u.branch_id
	LEFT JOIN sc_users.company c ON c.id = u.company_id
	LEFT JOIN sc_users.department d ON d.id = p.department_id
	LEFT JOIN sc_users.user_profile up ON up.user_id = u.id
	WHERE u.id = $1 LIMIT 1`, id)
	return user, err
}

// All ...
func (m UserModel) All() (user []interType.User, err error) {
	qs := `SELECT 
		u.id, 
		u.name, 
		u.email, 
		u.password,
		CASE 
			WHEN r.id IS NOT NULL THEN jsonb_build_object('id', r.id, 'name', r.name) 
			ELSE NULL 
		END AS role,
		CASE 
			WHEN c.id IS NOT NULL THEN jsonb_build_object('id', c.id, 'name', c.name) 
			ELSE NULL 
		END AS company,
		CASE 
			WHEN b.id IS NOT NULL THEN jsonb_build_object('id', b.id, 'name', b.name, 'address', b.address, 'contact', b.contact) 
			ELSE NULL 
		END AS branch, 
		CASE 
			WHEN d.id IS NOT NULL THEN jsonb_build_object('id', d.id, 'name', d.name) 
			ELSE NULL 
		END AS department,
		CASE 
			WHEN p.id IS NOT NULL THEN jsonb_build_object('id', p.id, 'name', p.name) 
			ELSE NULL 
		END AS position,
		CASE
			WHEN up.id IS NOT NULL THEN jsonb_build_object(
				'full_name', up.full_name, 
				'birth_date', up.birth_date, 
				'birth_place', up.birth_place, 
				'address', up.address, 
				'phone_number', up.phone_number, 
				'gender', up.gender, 
				'photo', up.photo
			)
			ELSE NULL
		END AS profile,
		u.active,
		u.created_at, 
		u.updated_at
	FROM sc_users.users u
	LEFT JOIN sc_users.role r ON r.id = u.role_id
	LEFT JOIN sc_users.position p ON p.id = u.position_id
	LEFT JOIN sc_users.branch b ON b.id = u.branch_id
	LEFT JOIN sc_users.company c ON c.id = u.company_id
	LEFT JOIN sc_users.department d ON d.id = p.department_id
	LEFT JOIN sc_users.user_profile up ON up.user_id = u.id
	ORDER BY u.created_at DESC`

	_, err = db.GetDB().Select(&user, qs)
	return user, err
}
