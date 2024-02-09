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
		fmt.Println("Requesting /gen from ", r.RemoteAddr)
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			fmt.Println("Invalid method")
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
		fmt.Println("Requesting ", r.RequestURI, " from ", r.RemoteAddr)
		uuid := r.URL.Path[1:] // Get the UUID from the path
		urls, exists := redirects[uuid]
		if !exists || len(urls) == 0 {
			http.NotFound(w, r)
			fmt.Println("UUID not found")
			return
		}

		selectedUrl := urls[rand.Intn(len(urls))]

		http.Redirect(w, r, selectedUrl, http.StatusFound)
	})

	// welcome url
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		//log the request
		fmt.Println("Requesting /welcome from ", r.RemoteAddr)
		//return welcome message
		fmt.Fprintf(w, "Welcome to Traverse!\n")
	})

	// Start the HTTP server
	output := "Server is running on port 8080"
	fmt.Println(output)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
