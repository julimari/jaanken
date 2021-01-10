package main

import (
	"fmt"
	"net/http"
	"time"
	"https://github.com/julimari/jaanken/models" // 自分のパッケージ
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "サーバーからのお返事です！")

	//呼び出される際に『w』と『r』とうい名前のデータを渡す
	//wはユーザーに文字列の送信
	//rはユーザーに関する情報が入っている。どのURLか、IPアドレスはなど
}

func main() {
	muxHTTP := http.NewServeMux()
	//httpパッケージのNewServerMuxを呼び出して、返ってきたデータをmuxHTTPと言う変数に格納
	muxHTTP.HandleFunc("/", handler)
	//serveMuxはhandleFuncと言う関数を使うことで『/』と言うアドレスにアクセスがあったらhandlerを呼ぶように設定している

	ServerHTTP := &http.Server{
		Addr:           ":8080",
		Handler:        muxHTTP,
		ReadTimeout:    10 * time.Second, //１０秒でアクセス
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //データの最大数
	}

	err := ServerHTTP.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
