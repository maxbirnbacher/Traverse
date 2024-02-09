package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"github.com/google/uuid"
	"github.com/go-faker/faker/v4"
)

var redirects = make(map[string][]string)

func generateUUID() string {
	return uuid.New().String()
}

// generate a security token
func generateToken() string {
	return uuid.New().String()
}

var token = generateToken()

func main() {
	// Redirect url
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requesting ", r.RequestURI, " from ", r.RemoteAddr)
		requested_path := r.URL.Path[1:] // Get the requested path from the path
		urls, exists := redirects[uuid]
		if !exists || len(urls) == 0 {
			http.NotFound(w, r)
			fmt.Println("Path not found")
			return
		}

		selectedUrl := urls[rand.Intn(len(urls))]
		fmt.Println("Redirecting request to ", selectedUrl)
		http.Redirect(w, r, selectedUrl, http.StatusFound)
	})

	// welcome url
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		//log the request
		fmt.Println("Requesting /welcome from ", r.RemoteAddr)
		//return welcome message
		fmt.Fprintf(w, "Welcome to Traverse!\n")
		//create a new UUID
		uuid := generateUUID()
		//add the UUID to the redirects map
		redirects[uuid] = []string{"https://www.google.com"}
		//return the UUID
		fmt.Println("Generated URL: /", uuid, " to ", "https://www.google.com")
		fmt.Fprintf(w, "Generated URL: /%s\n", uuid)
	})

	// add a redirect url
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		//log the request
		fmt.Println("Requesting /add from ", r.RemoteAddr)
		//parse the query parameters
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		//generate a fake URI path with faker
		path := faker.Internet().uri_path()
		//get the URL from the query parameters
		url := r.Form.Get("url")
		//add the URL to the redirects map
		redirects[uuid] = append(redirects[path], url)
		//return the URL
		fmt.Println("Generated URL: /", path, " to ", url)
		fmt.Fprintf(w, "Generated URL: /%s to %s\n", path, url)
	})

	// Start the HTTP server
	output := "Server is running on port 8080"
	fmt.Println(output)
	fmt.Println("Use the security token to create a initial url: ", token)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
