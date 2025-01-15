package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (u *UrlShortener) OriginUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//make sure is post
	if r.Method != "POST" {
		http.Error(w, "Wrong Method", http.StatusInternalServerError)
		return
	}

	shortUrl := r.FormValue("short_url")
	ShortKey := strings.TrimPrefix(shortUrl, u.baseURL)

	if originalURL, exist := u.urlMap[ShortKey]; exist {
		//long url already save return same url and key
		response := map[string]string{"url": originalURL, "short_url": shortUrl}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return

	}
	http.Error(w, "URL Not Found", http.StatusBadRequest)
	return

}

// func OriginUrl(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	//make sure is post
// 	if r.Method != "POST" {
// 		http.Error(w, "Wrong Method", http.StatusInternalServerError)
// 		return
// 	}

// 	baseURL := os.Getenv("BASE_URL")

// 	if baseURL == "" {
// 		baseURL = "http://localhost:8080/short/"
// 	}

// 	shortUrl := r.FormValue("short_url")
// 	ShortKey := strings.TrimPrefix(shortUrl, baseURL)

// 	if originalURL, exist := urlMap[ShortKey]; exist {
// 		//long url already save return same url and key
// 		response := map[string]string{"url": originalURL, "short_url": shortUrl}
// 		if err := json.NewEncoder(w).Encode(response); err != nil {
// 			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 			return
// 		}
// 		return

// 	}
// 	http.Error(w, "URL Not Found", http.StatusBadRequest)
// 	return

// }
