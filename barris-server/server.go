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
	Name          string `json:"name"`
	Url           string `json:"url"`
	TelegramToken string `json:"telegramToken"`
	Email         string `json:"email"`
}

var (
	name string
	url  string
)

type Barris struct {
	Barris []Barri `json:"barris"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to barri server!")
}

func (s *Server) addBarri(c echo.Context) error {
	b := &Barri{}
	if err := c.Bind(b); err != nil {
		return err
	}
	fmt.Println("Add Barri: ", b.Name, " Url: ", b.Url, " b.TelegramToken: ", b.TelegramToken, " Email: ", b.Email)
	addBarriToDB(b.Name, b.Url, b.TelegramToken, b.Email, s.DB)
	return c.JSON(http.StatusCreated, b)
}

func addBarriToDB(name string, url string, telegramToken string, email string, db *sql.DB) {
	sqlStatement := "INSERT INTO barris (name, url, telegram_token, admin) VALUES ($1, $2, $3, $4)"
	res, err := db.Query(sqlStatement, name, url, telegramToken, email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func (s *Server) getBarris(c echo.Context) error {
	fmt.Println("Funcion getBarris")
	email := c.QueryParam("email")
	fmt.Println("email getBarris: ", email)
	sqlStatement := "SELECT name, url, telegram_token FROM barris"
	if email != "" {
		sqlStatement = "SELECT name, url, telegram_token FROM barris WHERE admin = '" + email + "'"
	}

	rows, err := s.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Barris{}

	for rows.Next() {
		aBarri := Barri{}
		err := rows.Scan(&aBarri.Name, &aBarri.Url, &aBarri.TelegramToken)
		if err != nil {
			log.Fatal(err)
		}
		result.Barris = append(result.Barris, aBarri)
		fmt.Println(result)
	}
	return c.JSON(http.StatusCreated, result)
}

func (s *Server) updateBarri(c echo.Context) error {
	b := &Barri{}
	if err := c.Bind(b); err != nil {
		return err
	}
	sqlStatement := "UPDATE barris SET url=$1, telegram_token=$2 WHERE name=$2;"
	res, err := s.DB.Query(sqlStatement, b.Url, b.TelegramToken, b.Name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	return c.JSON(http.StatusCreated, b)
}

func (s *Server) updateBarriToken(c echo.Context) error {
	b := &Barri{}
	if err := c.Bind(b); err != nil {
		return err
	}
	fmt.Println("Update Barri: ", b.Name, " b.TelegramToken: ", b.TelegramToken)
	sqlStatement := "UPDATE barris SET telegram_token=$1 WHERE name=$2;"
	res, err := s.DB.Query(sqlStatement, b.TelegramToken, b.Name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	return c.JSON(http.StatusCreated, b)
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
	e.POST("/updateBarri", server.updateBarriToken)
	e.POST("/updateBarriToken", server.updateBarriToken)
	e.Logger.Fatal(e.Start(":1323"))
}
