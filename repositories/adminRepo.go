package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	databases "test_mekar_api/database"
	"test_mekar_api/models"
	"test_mekar_api/utils"
)


func RegisterAdmin(username string, password string) (models.ResponeMessage, error) {

	var res models.ResponeMessage
	hashPass,err := utils.HashPass(password)
	con := databases.CreateCon()

	sql := utils.REGISTER_ADMIN

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(username,hashPass)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"lastId": "done",
	}
	return res, nil

}


func LoginAdmin(username string, password string) (bool, error) {

	var obj models.LoginUser
	var pwd string

	con := databases.CreateCon()
	sqlDb := utils.LOGIN_ADMIN

	err := con.QueryRow(sqlDb, username).Scan(&obj.Id_user, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := utils.CheckPassword(password, pwd)
	if !match {
		fmt.Println("password tidak saama")
		return false, err
	}

	return true, err

}