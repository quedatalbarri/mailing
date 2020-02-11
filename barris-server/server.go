package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "quedatalbarri"
)

type barri struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome a barri server!")
}

func addBarri(c echo.Context) error {
	fmt.Println("Funcion addBarri")
	b := &barri{}
	if err := c.Bind(b); err != nil {
		return err
	}
	fmt.Println("Barri: ", b.Name, " Url: ", b.Url)
	addBarriToDB(b.Name)
	return c.JSON(http.StatusCreated, b)
}

func addBarriToDB(name string) {
	//connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	sqlStatement := "INSERT INTO barris (name, url)VALUES ($1, $2)"
	res, err := db.Query(sqlStatement, name, "b.Url")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

// func connectToDatabase() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("DB Connected...")
// 	}
// }

func main() {
	e := echo.New()

	// CORS restricted- Allows requests
	e.Use(middleware.CORS())

	//ROUTES
	e.GET("/", hello)
	e.POST("/addBarri", addBarri)

	e.Logger.Fatal(e.Start(":1323"))
}
