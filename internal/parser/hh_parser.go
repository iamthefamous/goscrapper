package parser

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/iamthefamous/goscrapper/internal/models"
)

func ParseHN(doc *goquery.Document) []models.Article {

	var articles []models.Article

	doc.Find(".titleline a").Each(func(i int, s *goquery.Selection) {

		title := s.Text()
		link, _ := s.Attr("href")

		articles = append(articles, models.Article{
			Title:     title,
			Link:      link,
			ScrapedAt: time.Now(),
		})
	})

	return articles
}
