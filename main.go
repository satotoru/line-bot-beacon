package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	var addr = flag.String("addr", ":8000", "アプリケーションのアドレス")

	http.HandleFunc("/line", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("error")
		}
		decoded, err := base64.StdEncoding.DecodeString(r.Header.Get("X-Line-Signature"))
		if err != nil {
			log.Println("error")
		}
		hash := hmac.New(sha256.New, []byte(os.Getenv("LINE_CHANNEL_SECRET")))
		hash.Write(body)
		if hmac.Equal(hash.Sum(nil), decoded) {
			log.Println("success")
		}
		// Compare decoded signature and `hash.Sum(nil)` by using `hmac.Equal`
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{}`))
	})
	// Webサーバーを開始
	log.Println("Webサーバを開始します。ポート:", *addr)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
