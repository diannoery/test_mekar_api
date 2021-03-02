package repositories

import (
	"net/http"
	databases "test_mekar_api/database"
	"test_mekar_api/models"
	"test_mekar_api/utils"
	"time"

	"github.com/google/uuid"
)

func FetchAllUser() (models.ResponeMessage, error) {
	var obj models.Users
	var arrObj []models.Users
	var res models.ResponeMessage

	con := databases.CreateCon()
	//sql := " select * from tb_users"
	sql := utils.ALL_USER

	rows, err := con.Query(sql)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.UserId, &obj.Nik, &obj.Nama, &obj.TanggalLahir, &obj.Pekerjaan.IDPekerjaan, &obj.Pekerjaan.Pekerjaan, &obj.PendidikanTerakhir.IdPendidikan, &obj.PendidikanTerakhir.Pendidikan, &obj.UserStatus, &obj.CreatedDate, &obj.UpdatedDate)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func FetchAllUserPage(i int , i2 int) ([]models.Users, error) {

	IndexFirst := (i * i2)-i2
	
	var obj models.Users
	var arrObj []models.Users
	var res models.ResponeMessage

	con := databases.CreateCon()
	
	sql := utils.ALL_USER_PAGE
	stmt, err := con.Prepare(sql)

	rows, err := stmt.Query(IndexFirst,i2)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.UserId, &obj.Nik, &obj.Nama, &obj.TanggalLahir, &obj.Pekerjaan.IDPekerjaan, &obj.Pekerjaan.Pekerjaan, &obj.PendidikanTerakhir.IdPendidikan, &obj.PendidikanTerakhir.Pendidikan, &obj.UserStatus, &obj.CreatedDate, &obj.UpdatedDate)

		if err != nil {
			return nil, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj
	

	return arrObj, nil
}

func DeleteUser(id string) (models.ResponeMessage, error) {
	
	var res models.ResponeMessage
	con := databases.CreateCon()

	sql := utils.DELETE_USER
	stmt, err := con.Prepare(sql)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	rowaffeted, err := result.RowsAffected()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowaffeted": rowaffeted,
	}
	return res, nil
}
func GetUserById(id string) (models.ResponeMessage, error) {
	var User models.Users
	var res models.ResponeMessage
	con := databases.CreateCon()

	sql := utils.GET_USER_BY_ID

	stmt, err := con.Prepare(sql)

	if err != nil {
		return res, err
	}

	err = stmt.QueryRow(id).Scan(&User.UserId, &User.Nik, &User.Nama, &User.TanggalLahir, &User.Pekerjaan.IDPekerjaan, &User.Pekerjaan.Pekerjaan, &User.PendidikanTerakhir.IdPendidikan, &User.PendidikanTerakhir.Pendidikan, &User.UserStatus, &User.CreatedDate, &User.UpdatedDate)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "User By Id"
	res.Data = User
	return res, nil
}
func CreateUser(nik string, nama string, TglLahir string, pekerjaan string, pendidikan string) (models.ResponeMessage, error) {
	var time = time.Now()
	var userId = uuid.New().String()
	var res models.ResponeMessage
	con := databases.CreateCon()

	sql := utils.CREATE_USERS

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(userId, nik, nama, TglLahir, pekerjaan, pendidikan, 1, time, time)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"lastId": nik,
	}
	return res, nil
}

func UpdateUser(id string, nik string, nama string, TglLahir string, pekerjaan string, pendidikan string) (models.ResponeMessage, error) {
	var time = time.Now()
	var res models.ResponeMessage
	con := databases.CreateCon()

	sql := utils.UPDATE_USER

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nik, nama, TglLahir, pekerjaan, pendidikan, 1, time, id)
	if err != nil {
		return res, err
	}

	rowaffeted, err := result.RowsAffected()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowaffeted": rowaffeted,
	}
	return res, nil
}


//pekerjaan
func FetchPekerjaan() (models.ResponeMessage, error) {
	var obj models.Pekerjaan
	var arrObj []models.Pekerjaan
	var res models.ResponeMessage

	con := databases.CreateCon()
	
	sql := utils.PEKERJAN

	rows, err := con.Query(sql)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IDPekerjaan, &obj.Pekerjaan)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}


//PENDIDIKAN
func FetchPendidikan() (models.ResponeMessage, error) {
	var obj models.Pendidikan
	var arrObj []models.Pendidikan
	var res models.ResponeMessage

	con := databases.CreateCon()
	
	sql := utils.PENDIDIKAN

	rows, err := con.Query(sql)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdPendidikan, &obj.Pendidikan)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}


func  CountUser() (int, error) {
	var totalData int
	con := databases.CreateCon()
	sql := utils.COUNT_DATA_USER
	stmt, err := con.Prepare(sql)
	if err != nil {
		return totalData, err
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&totalData)
	if err != nil {
		return totalData, err
	}
	return totalData, nil
}