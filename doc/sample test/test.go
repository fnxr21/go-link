package test

import (
	"fmt"
	"net/http"

	"github.com/fnxr21/go-link/pkg"
)

type URLShortener struct {
    urls map[string]string
}


func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    originalURL := r.FormValue("url")
    if originalURL == "" {
        http.Error(w, "URL parameter is missing", http.StatusBadRequest)
        return
    }

    // Generate a unique shortened key for the original URL
    shortKey := pkg.GenerateHexKey()
    us.urls[shortKey] = originalURL

    // Construct the full shortened URL
    shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

    // Render the HTML response with the shortened URL
    w.Header().Set("Content-Type", "text/html")
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
}




func (us *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
    shortKey := r.URL.Path[len("/short/"):]
    if shortKey == "" {
        http.Error(w, "Shortened key is missing", http.StatusBadRequest)
        return
    }

    // Retrieve the original URL from the `urls` map using the shortened key
    originalURL, found := us.urls[shortKey]
    if !found {
        http.Error(w, "Shortened key not found", http.StatusNotFound)
        return
    }

    // Redirect the user to the original URL
    http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}




func MainTest() {
    shortener := &URLShortener{
        urls: make(map[string]string),
    }

    http.HandleFunc("/test/shorten", shortener.HandleShorten)
    http.HandleFunc("/tes/short/", shortener.HandleRedirect)

    fmt.Println("URL Shortener is running on :8080")
    http.ListenAndServe(":8080", nil)
}