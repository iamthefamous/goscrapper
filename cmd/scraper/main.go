package main

import (
	"fmt"

	"log"

	"github.com/iamthefamous/goscrapper/internal/scraper"
	"github.com/iamthefamous/goscrapper/internal/storage"
)

func main() {

	collection, err := storage.Connect()
	if err != nil {
		log.Fatal(err)
	}

	url := "https://news.ycombinator.com/news?p=1"

	articles, err := scraper.ScrapePage(url)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.SaveArticles(collection, articles)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Saved", len(articles), "articles")
}
