package handler

import (
	// "fmt"
	"net/http"
	"strings"
)

func (u *UrlShortener) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// extract the short URL key by trimming the "/short/"  with prefix 
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")

	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}
	// retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := u.urlMap[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// redirect the user to the original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)

}

// func Redirect(w http.ResponseWriter, r *http.Request) {
// 	shortKey := r.URL.Path[len("/short/"):]

// 	if shortKey == "" {
// 		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
// 		return
// 	}

// 	for key, value := range urlMap {
// 		fmt.Printf("Key: %s, Value: %d\n", key, value)
// 	}
// 	// Retrieve the original URL from the `urls` map using the shortened key
// 	originalURL, found := urlMap[shortKey]
// 	if !found {
// 		http.Error(w, "Shortened key not found", http.StatusNotFound)
// 		return
// 	}

// 	// Redirect the user to the original URL
// 	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)

// }
