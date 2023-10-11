package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"
)

var rwmutex sync.RWMutex
var respData Cryptos
var indexPageData IndexPageData
var indexPageContentBuf bytes.Buffer

//go:embed index.css
var indexCssFile string

type IndexPageData struct {
	LastUpdateTime string
	Cryptos        Cryptos
}

type Cryptos []struct {
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"current_price"`
}

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))

	go func() {
		for {
			resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false")
			if err != nil {
				fmt.Println("Coin Gecko API error")
			} else if resp.StatusCode != http.StatusOK {
				fmt.Println("Coin Gecko API error")
			}

			if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
				fmt.Println("Coin Gecko API error")
			}

			resp.Body.Close()

			now := time.Now()

			lastUpdateTime := fmt.Sprintf("%d/%02d/%02d",
				now.Year(), now.Month(), now.Day(),
			)

			rwmutex.Lock()
			indexPageData = IndexPageData{
				LastUpdateTime: lastUpdateTime,
				Cryptos:        respData,
			}
			rwmutex.Unlock()

			if err := tmpl.Execute(&indexPageContentBuf, indexPageData); err != nil {
				break
			}

			time.Sleep(time.Minute * 5)
		}
	}()

	http.HandleFunc("/index.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(indexCssFile))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(indexPageContentBuf.Bytes())
	})

	http.ListenAndServe(":8080", nil)
}
