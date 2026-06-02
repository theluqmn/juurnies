package main

import (
	"main/util"
	
	"github.com/labstack/echo/v4"
)

func main() {
	util.Clear()
	util.LogInfo("server is initialising")

	err := util.Init("./database.sqlite")
	if err != nil { util.LogError(err) }
	
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	util.LogSuccess("server is now online")
	e.Logger.Fatal(e.Start(":6969"))
}