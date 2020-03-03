package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	Domain            string         `json:"domain"`
	NameObj           sql.NullString `json:"nameObj"`
	Name              *string        `json:"name"`
	Url               string         `json:"url"`
	TelegramChannelId string         `json:"telegramChannelId"`
	Email             string         `json:"email"`
}

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
	fmt.Println("Add Barri: ", b.Domain, " ", b.Name, " Url: ", b.Url, " b.TelegramChannelId: ", b.TelegramChannelId, " Email: ", b.Email)
	//addBarriToDB(b.Name, b.Url, b.TelegramChannelId, b.Email, s.DB)
	sqlStatement := "INSERT INTO barris (domain, name, url, telegram_channel, admin) VALUES ($1, $2, $3, $4, $5)"
	res, err := s.DB.Query(sqlStatement, b.Domain, b.Name, b.Url, b.TelegramChannelId, b.Email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	return c.JSON(http.StatusCreated, b)
}

// func addBarriToDB(name string, url string, telegramChannelId string, email string, db *sql.DB) {
// 	sqlStatement := "INSERT INTO barris (name, url, telegram_channel, admin) VALUES ($1, $2, $3, $4)"
// 	res, err := db.Query(sqlStatement, name, url, telegramChannelId, email)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 	}
// }

func (s *Server) getBarris(c echo.Context) error {
	email := c.QueryParam("email")
	fmt.Println("getBarris del usuario con email : ", email)
	sqlStatement := "SELECT domain, name, url, telegram_channel FROM barris"
	if email != "" {
		sqlStatement = "SELECT domain, name, url, telegram_channel FROM barris WHERE admin = '" + email + "'"
	}

	rows, err := s.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Barris{}

	for rows.Next() {
		aBarri := Barri{}
		err := rows.Scan(&aBarri.Domain, &aBarri.NameObj, &aBarri.Url, &aBarri.TelegramChannelId)
		if err != nil {
			log.Fatal(err)
		}
		aBarri.Name = &aBarri.NameObj.String

		result.Barris = append(result.Barris, aBarri)

	}
	return c.JSON(http.StatusCreated, result)
}

func (s *Server) updateBarri(c echo.Context) error {
	b := &Barri{}
	if err := c.Bind(b); err != nil {
		return err
	}
	sqlStatement := "UPDATE barris SET url=$1, telegram_channel=$2 WHERE domain=$2;"
	res, err := s.DB.Query(sqlStatement, b.Url, b.TelegramChannelId, b.Domain)
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

type MessageSent struct {
	Ok   bool `json:"ok"`
	Date string
	Text string
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
	barriDomain := c.Param("barri")
	//barriName := c.QueryParam("barri")
	// if barriName == "" {
	// 	return
	// }
	sqlStatement := "SELECT domain, telegram_channel FROM barris WHERE domain = '" + barriDomain + "'"

	rows, err := s.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Barris{}

	for rows.Next() {
		aBarri := Barri{}
		err := rows.Scan(&aBarri.Domain, &aBarri.TelegramChannelId)
		if err != nil {
			log.Fatal(err)
		}
		result.Barris = append(result.Barris, aBarri)
		fmt.Println(result)
	}
	return c.JSON(http.StatusCreated, result)
}

type ChannelMessage struct {
	Text  string `json:"text"`
	Barri string `json:"barri"`
}

func sendTelegramMessage(c echo.Context) error {
	channelName := c.Param("channel")
	msg := new(ChannelMessage)
	err1 := c.Bind(msg)
	if err1 != nil {
		fmt.Println("error")
	}
	fmt.Println(msg)
	var sendMessageUrl = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/sendMessage?chat_id=@" + channelName + "&text=" + url.QueryEscape(msg.Text) + "&parse_mode=Markdown"

	resp, err := http.Get(sendMessageUrl)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var messageSent MessageSent
	err = decoder.Decode(&messageSent)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("ok: ", messageSent.Ok, ", text: ", messageSent.Text)
	return c.JSON(http.StatusCreated, messageSent)
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
	e.GET("/barris", server.getBarris)
	e.POST("/barris", server.addBarri)
	//e.POST("/updateBarri", server.updateBarri)
	e.PUT("/barris/:barri", server.updateBarri)

	e.GET("/barris/:barri/channel", server.getBarriChannel)

	e.GET("/getChatMember/:channel", getChatMember)
	e.POST("/sendTelegramMessage/:channel", sendTelegramMessage)
	e.Logger.Fatal(e.Start(":1323"))
}
