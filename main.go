package main

import (
	"log"
	"net/http"

	
	"gemini-go/handlers"
	"gemini-go/lib"
)

func main() {
	utils.LoadEnv()

	http.HandleFunc("/ai", handlers.AiPromptHandler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
