package models

import "time"

type Article struct {
	Title     string    `bson:"title"`
	Link      string    `bson:"url"`
	ScrapedAt time.Time `bson:"scraped_at"`
}
