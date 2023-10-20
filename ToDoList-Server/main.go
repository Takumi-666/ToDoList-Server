package main

import (
	"github.com/Takumi-666/ToDoList-Server/model"
	"github.com/Takumi-666/ToDoList-Server/router"
    "github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
    // Goalテーブルを作成
    model.CreateTable(model.GetDBInstance())

    e.Start(":8000")
	
	
}