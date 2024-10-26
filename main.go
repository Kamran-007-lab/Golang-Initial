package main

import (
    "server/config"
    "server/routes"
    "github.com/labstack/echo/v4"
)

func main() {
    config.InitDB()
    e := echo.New()
    routes.SetupRoutes(e)
    e.Logger.Fatal(e.Start(":8080"))
}
