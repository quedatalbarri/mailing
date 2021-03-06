package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello, Laura")
	// Open our jsonFile
	jsonFile, err := os.Open("signups.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened signups.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to quedat al barri!")
	})
	e.GET("/:barri", func(c echo.Context) error {
		barri := c.Param("barri")
		url := result[barri]
		fmt.Println(url)
		return c.Redirect(http.StatusMovedPermanently, url)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
