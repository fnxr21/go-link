package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	// "strings"
)

func OriginUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//make sure is post
	if r.Method != "POST" {
		http.Error(w, "Wrong Method", http.StatusInternalServerError)
		return
	}
	shortUrl := r.FormValue("short_url")
	prefix := "http://localhost:8080/short/"

	ShortKey := strings.TrimPrefix(shortUrl, prefix)
	fmt.Println("Using TrimPrefix:", ShortKey)

	//lock for race condition
	//sample if we have same value want to change thats should be wait until this function done
	//defer make sure thats funtion will done

	mapMutex.Lock()
	defer mapMutex.Unlock()
	// port if exist use default port 8080
	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		baseURL = "http://localhost:8080/short/"
	}

	// // is more fast than use loop for scaning by row
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
