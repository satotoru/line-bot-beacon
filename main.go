package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	bot, err := linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Println("error")
	}

	http.HandleFunc("/line", func(w http.ResponseWriter, r *http.Request) {
		events, err := bot.ParseRequest(r)
		if err != nil {
			panic(err)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				_, err := bot.ReplyMessage(event.ReplyToken, event.Message).Do()
				if err != nil {
					panic(err)
				}
			} else if event.Type == linebot.EventTypeBeacon {
				_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ビーコンに近づきました")).Do()
				if err != nil {
					panic(err)
				}
			}
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{}`))
	})
	// Webサーバーを開始
	log.Println("Webサーバを開始します。")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
