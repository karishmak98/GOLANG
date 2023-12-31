package main

import (
	"PRACTICE/db"
	_ "PRACTICE/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

func main() {
	db.CreateCon()
	db.RedisCon()
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8000"))

}
