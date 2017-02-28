package main

import (
	"goApi1/dao"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	dao.InitDB("mssql", "driver={sql server};Server=182.92.65.253;port=1433;Database=SampleDB;user id=fruit;password=Eland123;")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello api")
	})
	v1 := e.Group("/v1")
	v1.GET("/fruits", GetAll)

	e.Start(":5000")
}

func GetAll(c echo.Context) error {
	fruitDao := dao.FruitDao{}
	if fruits, err := fruitDao.GetAll(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, fruits)
	}
}
