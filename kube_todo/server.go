package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" // Updated import path
	"github.com/joho/godotenv"
)

type TodoHomePageData struct {
	PageTitle string
	MainTitle string
	Todos     []string
}

func main() {

	//parse template
	tmpl := template.Must(template.ParseFiles("layout.html"))

	//Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Create a new router
	router := mux.NewRouter()

	// Handle requests
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeData := TodoHomePageData{
			PageTitle: "KUBE TODO",
			MainTitle: "welcome To Kube ToDo APP",
			Todos:     []string{"First Task"},
		}

		tmpl.Execute(w, homeData)
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("Starting server on port %s...", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
