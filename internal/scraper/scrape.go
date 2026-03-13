package scraper

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/iamthefamous/goscrapper/internal/models"
	"github.com/iamthefamous/goscrapper/internal/parser"
)

func ScrapePage(url string) ([]models.Article, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	articles := parser.ParseHN(doc)

	return articles, nil
}
