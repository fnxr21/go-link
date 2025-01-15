package handler

import (
	"encoding/json"
	"net/http"
	"os"
	// "path"
	"sync"

	"github.com/fnxr21/go-link/pkg"
)



type UrlShortener struct {
	urlMap     map[string]string
	reverseMap map[string]string
	mutex      sync.RWMutex
	baseURL    string
}

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
	if r.Method != "POST" {
		http.Error(w, "Wrong Method", http.StatusInternalServerError)
		return
	}

	//use formvalue
	originalURL := r.FormValue("url")

	//lock c
	u.mutex.Lock()
	defer u.mutex.Unlock()

	// // is more fast than use loop for scaning by row
	if shortURL, exist := u.reverseMap[originalURL]; exist {
		//long url already save return same url and key
		response := map[string]string{"url": originalURL, "short_url": u.baseURL + shortURL}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return

	}

	var shortURL string
	for {
		shortURL = pkg.GenerateHexKey()
		if _, exists := u.urlMap[shortURL]; !exists {
			break
		}
	}

	// Store in the maps
	u.urlMap[shortURL] = originalURL
	u.reverseMap[originalURL] = shortURL

	response := map[string]string{"url": originalURL, "short_url": u.baseURL + shortURL}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	return

}


//is not good for computing
//  for shortURL, storedURL := range urlMap {
//     if storedURL == longURL {
//         response := map[string]string{"shortURL": baseURL + shortURL}
//         json.NewEncoder(w).Encode(response)
//         return
//     }
// }


// var (
// 	urlMap     = make(map[string]string)
// 	reverseMap = make(map[string]string)
// 	mapMutex   = sync.RWMutex{} // Mutex for safe concurrent access
// )


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
