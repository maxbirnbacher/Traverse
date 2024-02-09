package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
)

var redirects = make(map[string][]string)

func generateUUID() string {
	return uuid.New().String()
}

func main() {
	http.HandleFunc("/gen", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var urls []string
		err := json.NewDecoder(r.Body).Decode(&urls)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		uuid := generateUUID()
		redirects[uuid] = urls

		fmt.Fprintf(w, "Generated URL: /%s\n", uuid)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Path[1:] // Get the UUID from the path
		urls, exists := redirects[uuid]
		if !exists || len(urls) == 0 {
			http.NotFound(w, r)
			return
		}

		selectedUrl := urls[rand.Intn(len(urls))]

		http.Redirect(w, r, selectedUrl, http.StatusFound)
	})

	// welcome url
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Traverse")
	})

	// Start the HTTP server
	output := "Server is running on port 8080"
	fmt.Println(output)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
