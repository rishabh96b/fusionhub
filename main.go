package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	BackgroundColor string
	Message         string
}

func handler(w http.ResponseWriter, r *http.Request) {
	color := os.Getenv("FUSIONHUB_COLOR")
	message := fmt.Sprintf("Hello, %s!", color)

	data := PageData{
		BackgroundColor: color,
		Message:         message,
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("FUSIONHUB_PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Server started on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
