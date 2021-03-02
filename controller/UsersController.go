package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"test_mekar_api/models"
	"test_mekar_api/repositories"

	"github.com/labstack/echo/v4"
)

func FecthAll(c echo.Context) error {

	page :=c.Param("page")
	limit :=c.Param("limit")
	fmt.Println(page)
	fmt.Println(limit)

	result, err := repositories.FetchAllUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FecthAllPage(c echo.Context) error {

	page :=c.Param("page")
	limit :=c.Param("limit")
	
	PageInt,_:= strconv.Atoi(page)
	LimitInt, _ := strconv.Atoi(limit)

	result, err := repositories.FetchAllUserPage(PageInt,LimitInt)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	totalData, err := repositories.CountUser()
	if err != nil {
		fmt.Println("error")
	}

	totalPage := math.Ceil(float64(totalData)/float64(LimitInt))

	res := models.UserList{
		Users: result,
		Metadata: models.Pagination{
			CurrentPage: PageInt,
			FirstPage:   1,
			LastPage:    int(totalPage),
			TotalData:   totalData,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	data:= models.Users{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		fmt.Print("error")
	}

	
	nik := data.Nik
	nama :=data.Nama
	TglLahir := data.TanggalLahir
	pekerjaan := data.Pekerjaan.IDPekerjaan
	pendidikan := data.PendidikanTerakhir.IdPendidikan


	result, err := repositories.CreateUser(nik, nama, TglLahir, pekerjaan, pendidikan)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func UpdateUser(c echo.Context) error {

	data:= models.Users{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		fmt.Print("error")
	}

	
	
	nik := data.Nik
	nama :=data.Nama
	TglLahir := data.TanggalLahir
	pekerjaan := data.Pekerjaan.IDPekerjaan
	pendidikan := data.PendidikanTerakhir.IdPendidikan

	idUser := data.UserId

	result, err := repositories.UpdateUser(idUser, nik, nama, TglLahir, pekerjaan, pendidikan)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}

func DeleteUser(c echo.Context) error  {
	id :=c.Param("id")

	result, err := repositories.DeleteUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetUserById(c echo.Context) error {
	id := c.Param("id")

	result, err := repositories.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//pekerjaan
func DataPekerjaan(c echo.Context) error {
	result, err := repositories.FetchPekerjaan()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//pendidikan
func DataPendidikan(c echo.Context) error {
	result, err := repositories.FetchPendidikan()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}