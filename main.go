package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	var addr = flag.String("addr", ":8000", "アプリケーションのアドレス")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	// Webサーバーを開始
	log.Println("Webサーバを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
