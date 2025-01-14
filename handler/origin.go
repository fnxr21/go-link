package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func OriginUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//make sure is post
	if r.Method != "POST" {
		http.Error(w, "Wrong Method", http.StatusInternalServerError)
		return
	}

	PORT := os.Getenv("APP_PORT")

	if PORT == "" {
		baseURL = "http://localhost:8080/short/"
	}

	
	shortUrl := r.FormValue("short_url")
	ShortKey := strings.TrimPrefix(shortUrl, baseURL)

	// mapMutex.Lock()
	// mapMutex.Unlock()

	if originalURL, exist := urlMap[ShortKey]; exist {
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
