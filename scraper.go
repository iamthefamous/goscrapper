package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapePage(url string) ([]Article, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var articles []Article

	doc.Find(".titleline a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")

		articles = append(articles, Article{
			Title: title,
			Link:  link,
		})
	})

	return articles, nil
}
