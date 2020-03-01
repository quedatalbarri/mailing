package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	Name              string `json:"name"`
	Url               string `json:"url"`
	TelegramChannelId string `json:"telegramChannelId"`
	Email             string `json:"email"`
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
	fmt.Println("Add Barri: ", b.Name, " Url: ", b.Url, " b.TelegramChannelId: ", b.TelegramChannelId, " Email: ", b.Email)
	addBarriToDB(b.Name, b.Url, b.TelegramChannelId, b.Email, s.DB)
	return c.JSON(http.StatusCreated, b)
}

func addBarriToDB(name string, url string, telegramChannelId string, email string, db *sql.DB) {
	sqlStatement := "INSERT INTO barris (name, url, telegram_channel, admin) VALUES ($1, $2, $3, $4)"
	res, err := db.Query(sqlStatement, name, url, telegramChannelId, email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func (s *Server) getBarris(c echo.Context) error {
	email := c.QueryParam("email")
	fmt.Println("getBarris del usuario con email : ", email)
	sqlStatement := "SELECT name, url, telegram_channel FROM barris"
	if email != "" {
		sqlStatement = "SELECT name, url, telegram_channel FROM barris WHERE admin = '" + email + "'"
	}

	rows, err := s.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Barris{}

	for rows.Next() {
		aBarri := Barri{}
		err := rows.Scan(&aBarri.Name, &aBarri.Url, &aBarri.TelegramChannelId)
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
	sqlStatement := "UPDATE barris SET url=$1, telegram_channel=$2 WHERE name=$2;"
	res, err := s.DB.Query(sqlStatement, b.Url, b.TelegramChannelId, b.Name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	return c.JSON(http.StatusCreated, b)
}

type ChatMember struct {
	Ok     bool `json:"ok"`
	Result Result
}

type Result struct {
	Status string
}

func getChatMember(c echo.Context) error {
	//loading enviroment variables
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}
	channelName := c.Param("channel")
	//	var url = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/getChatMember?chat_id=@pruebaquedat&user_id=1043999002"
	var url = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/getChatMember?chat_id=@" + channelName + "&user_id=1043999002"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var member ChatMember
	err = decoder.Decode(&member)
	if err != nil {
		fmt.Println("error")
	}
	//return member
	fmt.Println("ok: ", member.Ok, ", status: ", member.Result.Status)
	return c.JSON(http.StatusCreated, member)
}

func (s *Server) getBarriChannel(c echo.Context) error {
	barriName := c.Param("barri")
	//barriName := c.QueryParam("barri")
	// if barriName == "" {
	// 	return
	// }
	sqlStatement := "SELECT name, telegram_channel FROM barris WHERE name = '" + barriName + "'"

	rows, err := s.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Barris{}

	for rows.Next() {
		aBarri := Barri{}
		err := rows.Scan(&aBarri.Name, &aBarri.TelegramChannelId)
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
	e.POST("/updateBarri", server.updateBarri)
	e.GET("/getChatMember/:channel", getChatMember)
	e.GET("/getBarriChannel/:barri", server.getBarriChannel)
	e.Logger.Fatal(e.Start(":1323"))
}
