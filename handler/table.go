package handler

import (
	"html/template"
	"net/http"
)

type UrlData struct {
	ShortURL string
	LongURL  string
	Reversed string
}

func (u *UrlShortener) CheckTable(w http.ResponseWriter, r *http.Request) {
	//mapMutex.RLock // Lock for reading maps safely
	// 	defer mapMutex.RUnlock // unlock
	u.mutex.RLock()
	defer u.mutex.RUnlock()

	// Prepare data for the template
	var data []UrlData

	for shortURL, longURL := range u.urlMap {
		data = append(data, UrlData{
			ShortURL: shortURL,
			LongURL:  longURL,
			Reversed: u.reverseMap[longURL], // Get the reversed value from reverseMap
		})
	}

	var tmpl = template.Must(template.New("table_url").ParseFiles("view/view.html"))

	tmpl.Execute(w, data)

	return
}
