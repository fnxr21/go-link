package handler

import (
	"fmt"
	"html/template"
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

	// 
	if shortURL, exist := reverseMap[originalURL]; exist {
		//long url already save return same url and key
		response := map[string]string{"shortURL": baseURL + shortURL}
		fmt.Println(response)
		fmt.Println("same url")
		return

	}

	fmt.Println(originalURL)
	fmt.Println(baseURL)

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
	shortenedURL := fmt.Sprintf(baseURL, shortURL)

	// response := map[string]string{"shortURL": baseURL + shortURL}
	// w.Header().Set("Content-Type", "application/json")
	// // if err := json.NewEncoder(w).Encode(response); err != nil {
	// // 	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	// // 	return
	// // }
	// fmt.Println("==")
	// fmt.Println(pkg.DecodeHexKey(shortURL))
	// fmt.Println("==")
	responseHTML := fmt.Sprintf(`
	<h2>URL Shortener</h2>
	<p>Original URL: %s</p>
	<p>Shortened URL: <a href="%s">%s</a></p>
	<form method="post" action="/shorten">
		<input type="text" name="url" placeholder="Enter a URL">
		<input type="submit" value="Shorten">
	</form>
`, originalURL, shortenedURL, shortenedURL)
	fmt.Fprintf(w, responseHTML)
	return

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	for key, value := range urlMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// Retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := urlMap[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// Redirect the user to the original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)

}

// func CheckTable(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("==")
// 	for key, val := range urlMap {
// 		fmt.Println(key, "  \t:", val)
// 	}

// 	fmt.Println("==")
// 	for key, val := range reverseMap {
// 		fmt.Println(key, "  \t:", val)
// 	}

// 	fmt.Println("==")
// }

type UrlData struct {
	ShortURL string
	LongURL  string
	Reversed string
}

func CheckTable(w http.ResponseWriter, r *http.Request) {
	// Populate urlMap and reverseMap with some example data
	mapMutex.Lock()
	urlMap["abc123"] = "https://www.example.com"
	reverseMap["https://www.example.com"] = "abc123"
	mapMutex.Unlock()

	// Prepare data for the template
	var data []UrlData
	mapMutex.RLock() // Lock for reading maps safely

	// Iterate through urlMap and reverseMap to create a list of data
	for shortURL, longURL := range urlMap {
		data = append(data, UrlData{
			ShortURL: shortURL,
			LongURL:  longURL,
			Reversed: reverseMap[longURL], // Get the reversed value from reverseMap
		})
	}
	mapMutex.RUnlock()

	// Parse and execute the template
	// tmpl, err := template.New("table_url").Parse("view/view.html")

	var tmpl = template.Must(template.New("table_url").ParseFiles("view/view.html"))
	// Execute the template with the populated data
	tmpl.Execute(w, data)
	// if err != nil {
	// 	fmt.Println("Error executing template:", err)
	// }
	return
}
