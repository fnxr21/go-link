package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"

	"github.com/fnxr21/go-link/pkg"
)

var (
	urlMap     = make(map[string]string)
	reverseMap = make(map[string]string)
	mapMutex   = sync.RWMutex{} // Mutex for safe concurrent access
	baseURL    = "http://localhost:3000/short/"
)



func Shorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//make sure is post
	if r.Method != "POST" {
		http.Error(w, "Wrong Method", http.StatusInternalServerError)
		return
	}

	//use formvalue
	originalURL := r.FormValue("url")

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

	// # Find the short URL and return it

	//is not good for compo
	//  for shortURL, storedURL := range urlMap {
		//     if storedURL == longURL {
			//         response := map[string]string{"shortURL": baseURL + shortURL}
			//         json.NewEncoder(w).Encode(response)
			//         return
			//     }
			// }

	// // is more fast than use loop for scaning by row  
	if shortURL, exist := reverseMap[originalURL]; exist {
		//long url already save return same url and key
		response := map[string]string{"url": originalURL, "short_url": baseURL + shortURL}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return

	}

	var shortURL string
	for {
		shortURL = pkg.GenerateHexKey()
		if _, exists := urlMap[shortURL]; !exists {
			break
		}
	}

	// Store in the maps
	urlMap[shortURL] = originalURL
	reverseMap[originalURL] = shortURL

	response := map[string]string{"url": originalURL, "short_url": baseURL + shortURL}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	return

}



// func Shorten(w http.ResponseWriter, r *http.Request) {

// 	//make sure is post
// 	if r.Method != "POST" {
// 		http.Error(w, "Wrong Method", http.StatusInternalServerError)
// 		return
// 	}
// 	//use formvalue
// 	originalURL := r.FormValue("url")

// 	//lock for race condition
// 	//sample if we have same value want to change thats should be wait until this function done
// 	//defer make sure thats funtion will done

// 	mapMutex.Lock()
// 	defer mapMutex.Unlock()
// 	// port if exist use default port 8080
// 	PORT := os.Getenv("APP_PORT")
// 	if PORT == "" {
// 		baseURL = "http://localhost:8080/short/"
// 	}

// 	//
// 	if shortURL, exist := reverseMap[originalURL]; exist {
// 		//long url already save return same url and key
// 		response := map[string]string{"shortURL": baseURL + shortURL}
// 		fmt.Println(response)
// 		fmt.Println("same url")
// 		return

// 	}

// 	fmt.Println(originalURL)
// 	fmt.Println(baseURL)

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
// 	shortenedURL := fmt.Sprintf(baseURL, shortURL)

// 	// response := map[string]string{"shortURL": baseURL + shortURL}
// 	// w.Header().Set("Content-Type", "application/json")
// 	// // if err := json.NewEncoder(w).Encode(response); err != nil {
// 	// // 	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 	// // 	return
// 	// // }
// 	// fmt.Println("==")
// 	// fmt.Println(pkg.DecodeHexKey(shortURL))
// 	// fmt.Println("==")
// 	responseHTML := fmt.Sprintf(`
// 	<h2>URL Shortener</h2>
// 	<p>Original URL: %s</p>
// 	<p>Shortened URL: <a href="%s">%s</a></p>
// 	<form method="post" action="/shorten">
// 		<input type="text" name="url" placeholder="Enter a URL">
// 		<input type="submit" value="Shorten">
// 	</form>
// `, originalURL, shortenedURL, shortenedURL)
// 	fmt.Fprintf(w, responseHTML)
// 	return

// }
