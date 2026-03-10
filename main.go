package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(page int) {
			defer wg.Done()

			url := fmt.Sprintf("https://news.ycombinator.com/news?p=%d", page)

			data, err := ScrapePage(url)
			if err != nil {
				log.Println("error scraping:", err)
				return
			}

			fmt.Printf("Scraped page %d: %d results\n\n", page, len(data))

			for _, article := range data {
				fmt.Printf("Title: %s\nLink: %s\n\n", article.Title, article.Link)
			}

		}(i)
	}

	wg.Wait()
}
