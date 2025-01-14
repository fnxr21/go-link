package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fnxr21/go-link/handler"
	"github.com/joho/godotenv"
)

func RunServer() {
	dotEnv()

	//route
	http.HandleFunc("/", handler.Shorten)
	http.HandleFunc("/short/", handler.Redirect)//redirect
	http.HandleFunc("/table", handler.CheckTable)//table
	http.HandleFunc("/original-url", handler.OriginUrl)//table

	//testing
	// http.HandleFunc("/first", handler.RouteIndexGet)
	// http.HandleFunc("/process", handler.RouteSubmitPost)

	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		PORT = "8080"
	}
	fmt.Println("Server started at http://localhost:" + PORT)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func dotEnv()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
