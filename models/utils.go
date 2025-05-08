package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"

	"strconv"

	"github.com/fatih/structs"
	"github.com/google/uuid"
)

// UserSessionInfo ...
type UserSessionInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// JSONRaw ...
type JSONRaw json.RawMessage

// Value ...
func (j JSONRaw) Value() (driver.Value, error) {
	byteArr := []byte(j)
	return driver.Value(byteArr), nil
}

// Scan ...
func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}
	err := json.Unmarshal(asBytes, &j)
	if err != nil {
		return error(errors.New("Scan could not unmarshal to []string"))
	}
	return nil
}

// MarshalJSON ...
func (j *JSONRaw) MarshalJSON() ([]byte, error) {
	return *j, nil
}

// UnmarshalJSON ...
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// DataList ....
type DataList struct {
	Data JSONRaw `db:"data" json:"data"`
	Meta JSONRaw `db:"meta" json:"meta"`
}

func FlexibleUpdate(table_name string, str interface{}, cond *interType.UpdateCond, ret string) (err error) {

	query := "UPDATE " + table_name + " SET "
	j := 0
	s := structs.New(str)
	m := s.Map()

	var values []interface{}
	for i := range m {
		if v := m[i]; v != "" && v != 0 && v != 0.0 && !IsArrayEmpty(v) && !CheckTypeAndIsNullUUID(v) /* you're missing a condition here */ {
			j++
			query = query + "\"" + s.Field(i).Tag("form") + "\"" + "=$" + strconv.Itoa(j) + ","
			values = append(values, v)
		}
	}
	query += "updated_at=now(),"

	// adding conditions
	if cond.Ids != "" {
		query = query[:len(query)-1] + " WHERE " + cond.Ids + "=$" + strconv.Itoa(j+1)
		values = append(values, cond.Vals)
	}

	// return values
	if ret != "" {
		query = query + " RETURNING " + ret
	}
	// execution db
	operation, err := db.GetDB().Exec(query, values...)
	if err != nil {
		return err
	}
	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}
func IsArrayEmpty(arr interface{}) bool {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		v := reflect.ValueOf(arr)
		return v.Len() == 0
	default:
		return false
	}
}
func FlexibleInsert(table_name string, str interface{}, ret string) (id uuid.UUID, err error) {

	query := "INSERT INTO " + table_name
	j := 0
	s := structs.New(str)
	m := s.Map()
	qcol := "("
	qval := ") VALUES ("
	var values []interface{}
	for i := range m {
		if v := m[i]; v != "" && v != 0 && v != 0.0 && !IsArrayEmpty(v) && !CheckTypeAndIsNullUUID(v) /* you're missing a condition here */ {
			j++
			qcol = qcol + "\"" + s.Field(i).Tag("form") + "\"" + ","
			qval = qval + "$" + strconv.Itoa(j) + ","
			values = append(values, v)
		}
	}
	query = query + qcol[:len(qcol)-1] + qval[:len(qval)-1] + ")"

	// return values
	if ret != "" {
		query = query + " RETURNING " + ret
	}
	// execution db
	err = db.GetDB().QueryRow(query, values...).Scan(&id)
	return id, err
}
func FlexibleInsertIdString(table_name string, str interface{}, ret string) (id string, err error) {

	query := "INSERT INTO " + table_name
	j := 0
	s := structs.New(str)
	m := s.Map()
	qcol := "("
	qval := ") VALUES ("
	var values []interface{}
	for i := range m {
		if v := m[i]; v != "" && v != 0 && v != 0.0 && !IsArrayEmpty(v) && !CheckTypeAndIsNullUUID(v) /* you're missing a condition here */ {
			j++
			qcol = qcol + "\"" + s.Field(i).Tag("form") + "\"" + ","
			qval = qval + "$" + strconv.Itoa(j) + ","
			values = append(values, v)
		}
	}
	query = query + qcol[:len(qcol)-1] + qval[:len(qval)-1] + ")"

	// return values
	if ret != "" {
		query = query + " RETURNING " + ret
	}
	// execution db
	err = db.GetDB().QueryRow(query, values...).Scan(&id)
	return id, err
}
func FlexibleInsertIdInt(table_name string, str interface{}, ret string) (id int64, err error) {

	query := "INSERT INTO " + table_name
	j := 0
	s := structs.New(str)
	m := s.Map()
	qcol := "("
	qval := ") VALUES ("
	var values []interface{}
	for i := range m {
		if v := m[i]; v != "" && v != 0 && v != 0.0 && !IsArrayEmpty(v) && !CheckTypeAndIsNullUUID(v) /* you're missing a condition here */ {
			j++
			qcol = qcol + "\"" + s.Field(i).Tag("form") + "\"" + ","
			qval = qval + "$" + strconv.Itoa(j) + ","
			values = append(values, v)
		}
	}
	query = query + qcol[:len(qcol)-1] + qval[:len(qval)-1] + ")"

	// return values
	if ret != "" {
		query = query + " RETURNING " + ret
	}
	// execution db
	err = db.GetDB().QueryRow(query, values...).Scan(&id)
	return id, err
}

func FlexibleDelete(table_name string, cond *interType.UpdateCond) (err error) {

	query := "DELETE FROM " + table_name
	j := 0

	var values []interface{}

	// delete conditions
	if cond.Ids != "" {
		query = query + " WHERE " + cond.Ids + "=$" + strconv.Itoa(j+1)
		values = append(values, cond.Vals)
	}

	// execution db
	operation, err := db.GetDB().Exec(query, values...)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("deleted 0 records")
	}

	return err
}

func LogActivity(activity *interType.LogActivity) {
	// Prepare statement untuk insert data ke dalam tabel
	qs := "INSERT INTO sc_users.log_activity (user_id, type, detail) VALUES ($1, $2, $3)"

	// Melakukan eksekusi perintah SQL untuk menambahkan aktivitas baru
	_, err := db.GetDB().Exec(qs, activity.UserID, activity.Type, activity.Detail)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Activity logged successfully!")
}

func GetTotalCount(tableName string, cond *interType.UpdateCond) (total int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	var values []interface{}
	// where conditions
	if cond != nil && cond.Ids != "" {
		query = query + " WHERE " + cond.Ids + "=$" + strconv.Itoa(1)
		values = append(values, cond.Vals)
	}
	// Query to fetch total count
	totalCount, err := db.GetDB().SelectInt(query, values...)

	if err != nil {
		return 0, err
	}
	return totalCount, err
}

func GetTotalCountMultyCond(tableName string, cond []*interType.UpdateCond) (total int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	var values []interface{}
	// where conditions
	if cond != nil && !IsArrayEmpty(cond) {
		qcol := " WHERE "
		j := 0
		for _, v := range cond {
			if v.Ids != "" && v.Vals != "" {
				j++
				qcol = qcol + "\"" + v.Ids + "\"" + "=$" + strconv.Itoa(j) + "AND"
				values = append(values, v.Vals)
			}
		}
		query = query + qcol[:len(qcol)-3]
	}

	// Query to fetch total count
	totalCount, err := db.GetDB().SelectInt(query, values...)

	if err != nil {
		return 0, err
	}
	return totalCount, err
}

func CheckTypeAndIsNullUUID(data interface{}) bool {
	dataType := reflect.TypeOf(data)
	if (dataType == reflect.TypeOf(uuid.UUID{})) {
		return data == (uuid.UUID{})
	} else {
		return false
	}
}

