package main

import (
	"fmt"
	"net/http"
)

"""
ハンドラ関数(コントローラ/モデル)
PrintfがResponseWriterを介して書き出しするので引数として必要
Requestにアクセスするにはポインタで, %sにこのURIからの情報が必要。(なんでポインタ?)
"""
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World,%s!", request.URL.Path[1:])
}

"""
必ず必要
ルートに対してはhandlerを起動する。
ポート8080を監視するサーバを起動。
"""
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

