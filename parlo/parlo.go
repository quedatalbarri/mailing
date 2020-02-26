package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Event struct {
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	HtmlLink    string `json:"link,omitempty"`
	DateTime    string `json:"datetime,omitempty"`
}

//type Week struct {
type Events struct {
	Events []Event `json:"events,omitempty"`
}

type EmailContent struct {
	Metadata map[string]string
	Events   []Event
}

func main() {

	//leer json:
	jsonFile, err := os.Open("events.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var events Events
	json.Unmarshal(byteValue, &events)
	var eventsText string = ""
	for i := 0; i < len(events.Events); i++ {
		if i > 0 {
			eventsText = eventsText + "\n\n"
		}
		e := events.Events[i]
		eventsText = eventsText + "*" + e.Summary + "*\n" + e.Description + "\n" + e.Location + "\n" + e.HtmlLink + e.DateTime
	}

	//loading enviroment variables
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}
	//telegram api:

	poller := &tb.LongPoller{Timeout: 10 * time.Second}
	b, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("TELEGRAM_TOKEN"),
		//URL: You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		Poller: poller,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})
	b.Handle(tb.OnText, func(m *tb.Message) {
		fmt.Println("OnText...")
		b.Send(m.Sender, eventsText, &tb.SendOptions{
			ParseMode: "Markdown"},
		)
	})
	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		fmt.Println("OnChannelPost...")
		b.Send(m.Chat, eventsText, &tb.SendOptions{
			ParseMode: "Markdown"},
		)
		//b.Send(m.Chat, "Hola, soy el bot")
	})
	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		fmt.Println("OnUserJoined...")
		//no funciona
	})
	b.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		fmt.Println("OnAddedToGroup ...")
		//no funciona OnAddedToGroup
	})
	//new_chat_members
	b.Handle(tb.OnCallback, func(m *tb.Message) {
		fmt.Println("OnCallback...")
	})
	b.Handle(tb.OnQuery, func(m *tb.Message) {
		fmt.Println("OnQuery ...")
	})

	b.Start()
}
