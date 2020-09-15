package main

import (
	"log"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

//bot, err := linebot.New("5d92fc910eb38c80d75c384834a8d00f", "rfd85z4vWhZxcTS1C8sN4snCy1Bj7sPrLZwN3//jKWkbnripHvNE3vSErSdzgq/gcOKyjGDarDIgBgJ8mBNe7A69twaUSz/gyt/U7aaPTNWzyH7MhPoy7Bb/9nlNKbY0KEjcj6VvpDFV1E/oUKbHKQdB04t89/1O/w1cDnyilFU=")

func main() {
	handler, err := httphandler.New("5d92fc910eb38c80d75c384834a8d00f", "rfd85z4vWhZxcTS1C8sN4snCy1Bj7sPrLZwN3//jKWkbnripHvNE3vSErSdzgq/gcOKyjGDarDIgBgJ8mBNe7A69twaUSz/gyt/U7aaPTNWzyH7MhPoy7Bb/9nlNKbY0KEjcj6VvpDFV1E/oUKbHKQdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "w" {
						time.Sleep(2 * time.Second)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ready")).Do()
					} else {
						if message.Text == "おはよう" {
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("สวัสดี")).Do()
						}
						if message.Text == "สวัสดี" {
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("おはよう")).Do()
						}
					}

					log.Println(message.Text)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("unknown?")).Do()

					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					// 	log.Print(err)
					// }
				case *linebot.StickerMessage:
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hello Sticker")).Do()
				}
			}
		}
	})
	http.Handle("/callback", handler)
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
