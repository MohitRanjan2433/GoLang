package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type URL struct{
	ID            string      `json:"id"`
	OriginalURL   string      `json:"original_url"`
	ShortenURL    string 	  `json:"short_url"`
	CreationDate  time.Time   `json:"creation_date"`
}

/*
	d976351 --->  {
			ID: "d976351",
			OrignalUrl: "https:github.com",
			ShortUrl: "d976351",
			CreationDate: time.Now()
				}

*/

var urlDB = make(map[string]URL)

func generateShortUrl(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data", data)
	hash := hex.EncodeToString(data)
	fmt.Println("EncodeToString", hash)
	fmt.Println("final String", hash[:8])
	return hash[:8]
}

func main(){
	fmt.Println("Starting Url Shortner...")
	OriginalURL := "https://github.com/RnjnMohit"
	generateShortUrl(OriginalURL)
}

