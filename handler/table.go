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

func CheckTable(w http.ResponseWriter, r *http.Request) {
	mapMutex.Lock()

	defer mapMutex.Unlock()

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

	tmpl.Execute(w, data)

	return
}
