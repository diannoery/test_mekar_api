package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test_mekar_api/middleware"
	"test_mekar_api/models"
	"test_mekar_api/repositories"

	"github.com/labstack/echo/v4"
)

func RegisterAdmin(c echo.Context) error {
	register := models.LoginUser{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&register)
	if err != nil {
		fmt.Print("error")
	}

	result, err := repositories.RegisterAdmin(register.Username, register.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func LoginAdmin(c echo.Context) error {
	login := models.LoginUser{}
	var res models.ResponeMessage
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&login)
	if err != nil {
		fmt.Print("error")
	}

	result, err := repositories.LoginAdmin(login.Username, login.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if !result {
		return echo.ErrUnauthorized
	}

	token,_ := middleware.CreateToken(login.Username)
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"token": token,
	}


	return c.JSON(http.StatusOK,res)
}
