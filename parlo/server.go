package main

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	poller := &tb.LongPoller{Timeout: 10 * time.Second}
	b, err := tb.NewBot(tb.Settings{
		Token: "",
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		//URL:    "http://195.129.111.17:8012",
		Poller: poller,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("", func(m *tb.Message) {
		fmt.Println("hola ...")
	})

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})
	b.Handle(tb.OnText, func(m *tb.Message) {
		fmt.Println("OnText...")
		//b.Send(m.Sender, "hello")
	})
	b.Handle(tb.OnCallback, func(m *tb.Message) {
		fmt.Println("OnCallback...")
		//b.Send(m.Sender, "hello")
	})
	b.Handle(tb.OnQuery, func(m *tb.Message) {
		fmt.Println("OnQuery ...")
		//b.Send(m.Sender, "hello")
	})
	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		fmt.Println("OnChannelPost...")
		b.Send(m.Chat, "Hola, soy el bot")
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
	b.Start()
}
