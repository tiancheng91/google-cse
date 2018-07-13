package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	cse "github.com/tiancheng91/google-cse"
)

func main() {
	// 请求频繁的化多加几个agent轮询试试
	agent := cse.New("008063188944472181627:xqha3yefaee", "zh_CN")

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		page, err := strconv.Atoi(query.Get("p"))
		res, err := agent.Query(query.Get("q"), int64(page), 20)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})
	log.Fatal(http.ListenAndServe(":18080", nil))
}
