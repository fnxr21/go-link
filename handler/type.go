package handler

import "sync"

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
