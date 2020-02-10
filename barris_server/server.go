package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	barri struct {
		Name string `json:"name"`
	}
)

var seq = 1
var xxx string

func createBarri(c echo.Context) error {
	fmt.Println("Funcion createBarri")
	u := &barri{
		Name: xxx,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println("Barri: ", u.Name)
	//TODO add barrio to a database
	return c.JSON(http.StatusCreated, u)
}

func main() {
	fmt.Println("Funcion main")

	e := echo.New()

	// CORS restricted
	// Allows requests
	e.Use(middleware.CORS())
	//ROUTES
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome a barri server!")
	})
	e.POST("/addBarri", createBarri)
	e.Logger.Fatal(e.Start(":1323"))
}
