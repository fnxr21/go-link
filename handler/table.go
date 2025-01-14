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
	// Populate urlMap and reverseMap with some example data
	mapMutex.Lock()
	// urlMap["abc123"] = "https://www.example.com"
	// reverseMap["https://www.example.com"] = "abc123"
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

