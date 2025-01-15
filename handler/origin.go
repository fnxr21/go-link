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

	// shortUrl := r.FormValue("short_url")

	var requestData RequestShortURL
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if URL is provided
	if requestData.ShortURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	ShortKey := strings.TrimPrefix(requestData.ShortURL, u.baseURL)

	if originalURL, exist := u.urlMap[ShortKey]; exist {
		//long url already save return same url and key
		// response := map[string]string{"url": originalURL, "short_url": requestData.ShortURL}
		response := ResponseURL{
			URL:      originalURL,
			ShortURL: requestData.ShortURL,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return

	}
	http.Error(w, "URL Not Found", http.StatusBadRequest)
	return

}
