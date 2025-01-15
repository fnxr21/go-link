package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/fnxr21/go-link/pkg"
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

func (u *UrlShortener) Shorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//make sure is post
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData RequestLongURL
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if URL is provided
	if requestData.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	//lock trancation
	// make sure thats not gonna make race connditon for making new url in same time
	u.mutex.Lock()
	defer u.mutex.Unlock()

	//is not good for computing is better acces direct
	//  for shortURL, storedURL := range urlMap {
	//     if storedURL == longURL {
	//         response := map[string]string{"shortURL": baseURL + shortURL}
	//         json.NewEncoder(w).Encode(response)
	//         return
	//     }
	// }

	// is more fast than use loop for scaning by row like above

	// Check if the original URL already has a shortened URL
	if shortURL, exist := u.reverseMap[requestData.URL]; exist {
		//long url already save return same url and key
		response := ResponseURL{
			URL:      requestData.URL,
			ShortURL: u.baseURL + shortURL,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return

	}

	// generate unique shortURL
	var shortURL string
	for {
		shortURL = pkg.GenerateHexKey()
		if _, exists := u.urlMap[shortURL]; !exists {
			break
		}
	}

	// Store in the hashmap
	u.urlMap[shortURL] = requestData.URL
	u.reverseMap[requestData.URL] = shortURL
	response := ResponseURL{
		URL:      requestData.URL,
		ShortURL: u.baseURL + shortURL,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	return

}

// type UrlShortener struct {
// 	urlMap     map[string]string
// 	reverseMap map[string]string
// 	mutex      sync.RWMutex
// 	baseURL    string
// }

// func Shorten(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	baseURL := os.Getenv("BASE_URL")
// 	if baseURL == "" {
// 		baseURL = "http://localhost:8080/short/"
// 	}

// 	//make sure is post
// 	if r.Method != "POST" {
// 		http.Error(w, "Wrong Method", http.StatusInternalServerError)
// 		return
// 	}

// 	//use formvalue
// 	originalURL := r.FormValue("url")

// 	//lock c
// 	mapMutex.Lock()
// 	defer mapMutex.Unlock()

// 	// // is more fast than use loop for scaning by row
// 	if shortURL, exist := reverseMap[originalURL]; exist {
// 		//long url already save return same url and key
// 		response := map[string]string{"url": originalURL, "short_url": baseURL + shortURL}
// 		if err := json.NewEncoder(w).Encode(response); err != nil {
// 			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 			return
// 		}
// 		return

// 	}

// 	var shortURL string
// 	for {
// 		shortURL = pkg.GenerateHexKey()
// 		if _, exists := urlMap[shortURL]; !exists {
// 			break
// 		}
// 	}

// 	// Store in the maps
// 	urlMap[shortURL] = originalURL
// 	reverseMap[originalURL] = shortURL

// 	response := map[string]string{"url": originalURL, "short_url": baseURL + shortURL}
// 	if err := json.NewEncoder(w).Encode(response); err != nil {
// 		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 		return
// 	}

// 	return

// }
