package models

import (
	"crud-go/db"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type Waifu struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Full_Name string `json:"full name" validate:"required"`
}

func FetchAllWaifu() (Response, error) {
	var obj Waifu
	var arrayObj []Waifu
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM waifu"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Name, &obj.Title, &obj.Full_Name, &obj.Id)
		if err != nil {
			return res, err
		}

		arrayObj = append(arrayObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayObj

	return res, nil
}

func StoreWaifu(name, title, full_name string) (Response, error) {
	var res Response

	v := validator.New()

	waip := Waifu{
		Name:      name,
		Title:     title,
		Full_Name: full_name,
	}

	err := v.Struct(waip)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT waifu (name, title, full_name) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, title, full_name)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{"last_inserted_id": lastInsertedId}

	return res, nil
}

func UpdateWaifu(name, title, full_name string, id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE waifu SET name = ?, title = ?, full_name = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, title, full_name, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteWaifu(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM waifu WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
