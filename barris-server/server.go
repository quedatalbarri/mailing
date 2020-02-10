package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type barri struct {
	Name string `json:"name"`
}

func hello(c.echo.Context) error {
	return c.String(http.StatusOK, "Welcome a barri server!")
}

func addBarri(c echo.Context) error {
	fmt.Println("Funcion addBarri")
	b := &barri{}
	if err := c.Bind(b); err != nil {
		return err
	}
	fmt.Println("Barri: ", b.Name)
	//TODO add barrio to a database
	return c.JSON(http.StatusCreated, b)
}

func main() {
	e := echo.New()

	// CORS restricted- Allows requests
	e.Use(middleware.CORS())
	
	//ROUTES
	e.GET("/", hello)
	e.POST("/addBarri", addBarri)

	e.Logger.Fatal(e.Start(":1323"))
}
