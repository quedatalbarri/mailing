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

type Barri struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var (
	name string
	url  string
)

type Barris struct {
	Barris []Barri `json:"barris"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome a barri server!")
}

func (s *Server) addBarri(c echo.Context) error {
	b := &Barri{}
	if err := c.Bind(b); err != nil {
		return err
	}
	fmt.Println("Add Barri: ", b.Name, " Url: ", b.Url)
	addBarriToDB(b.Name, b.Url, s.DB)
	return c.JSON(http.StatusCreated, b)
}

func addBarriToDB(name string, url string, db *sql.DB) {
	sqlStatement := "INSERT INTO barris (name, url)VALUES ($1, $2)"
	res, err := db.Query(sqlStatement, name, url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func (s *Server) getBarris(c echo.Context) error {
	fmt.Println("Funcion getBarris")
	sqlStatement := "SELECT name, url FROM barris"
	rows, err := s.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Barris{}

	for rows.Next() {
		aBarri := Barri{}
		err := rows.Scan(&aBarri.Name, &aBarri.Url)
		if err != nil {
			log.Fatal(err)
		}
		result.Barris = append(result.Barris, aBarri)
		fmt.Println(result)
	}
	return c.JSON(http.StatusCreated, result)
}

func connectToDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
	return db
}

type Server struct {
	DB *sql.DB
}

func main() {
	db := connectToDatabase()
	server := &Server{db}
	e := echo.New()

	// CORS restricted- Allows requests
	e.Use(middleware.CORS())

	//ROUTES
	e.GET("/", hello)
	e.GET("/getBarris", server.getBarris)
	e.POST("/addBarri", server.addBarri)

	e.Logger.Fatal(e.Start(":1323"))
}
