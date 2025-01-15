package handler

import (
	"os"
	"sync"
)

type (
	UrlShortener struct {
		urlMap     map[string]string
		reverseMap map[string]string
		mutex      sync.RWMutex
		baseURL    string
	}
	RequestLongURL struct {
		URL string `json:"url"`
	}
	RequestShortURL struct {
		ShortURL string `json:"short_url"`
	}
	ResponseURL struct {
		URL      string `json:"url"`
		ShortURL string `json:"short_url"`
	}
)


func NewUrlshortener() *UrlShortener {

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/short/"
	}

	return &UrlShortener{
		urlMap:     make(map[string]string),
		reverseMap: make(map[string]string),
		baseURL:    baseURL,
	}
}