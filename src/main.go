package main

import (
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

// generate a security token
func generateToken() string {
	return uuid.New().String()
}

var token = generateToken()

func main() {
	// Redirect url
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
		//get the UUID from the form
		uuid := generateUUID()
		//get the URL from the form
		url := r.FormValue("url")
		//add the URL to the redirects map
		redirects[uuid] = append(redirects[uuid], url)
		//return success message
		fmt.Fprintf(w, "Added URL: %s to /%s\n", url, uuid)
	})

	// Start the HTTP server
	output := "Server is running on port 8080"
	fmt.Println(output)
	fmt.Println("Use the security token to create a initial url: ", token)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
