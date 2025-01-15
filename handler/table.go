package handler

import (
	"html/template"
	"net/http"
)

type UrlData struct {
	Key      string
	LongURL  string
	ShortURL string
}

func (u *UrlShortener) CheckTable(w http.ResponseWriter, r *http.Request) {
	//mapMutex.RLock // Lock for reading maps safely
	// 	defer mapMutex.RUnlock  // unlock after this function done
	u.mutex.RLock()
	defer u.mutex.RUnlock()

	// Prepare data for the template
	var data []UrlData
	for keyURL, longURL := range u.urlMap {
		// populate data with short and long URLs, and the reversed key
		data = append(data, UrlData{
			Key:      keyURL,
			LongURL:  longURL,
			ShortURL: u.baseURL + keyURL,
			// Reversed: u.reverseMap[longURL], // get the reversed value from reverseMap
		})
	}

	var tmpl = template.Must(template.New("table_url").ParseFiles("view/view.html"))

	tmpl.Execute(w, data)

	return
}
