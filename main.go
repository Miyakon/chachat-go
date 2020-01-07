package main

import (
	"log"
	"net/http"
)

func main() {
	//ディレクトリを指定する
	fs := http.FileServer(http.Dir("static"))
	//ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示
	http.Handle("/", fs) 

	log.Println("Listening...")
	//3000ポートでサーバを立ち上げる
	http.ListenAndServe(":3000", nil)
}
