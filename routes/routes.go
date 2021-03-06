package routes

import (
	"net/http"
	"test_mekar_api/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "")
	})
	//user
	e.GET("/users", controller.FecthAll)
	e.GET("/users/:page/:limit", controller.FecthAllPage)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.GET("/users/:id", controller.GetUserById)

	//pekerjaan
	e.GET("/pekerjaan",controller.DataPekerjaan)

	//pendidikan
	e.GET("/pendidikan",controller.DataPendidikan)

	//login & register admin
	e.POST("/register",controller.RegisterAdmin)
	e.POST("/login",controller.LoginAdmin)



	return e
}
