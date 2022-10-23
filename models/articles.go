package models

import (
	"time"
)

type Time time.Time

type Article struct {
	Title       string   `json:"title"`
	Date        Time     `json:"date"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Author      string   `json:"author"`
	MusicId     string   `json:"musicId"`
	Path        string
	ShortUrl    string
	Category    string
}
