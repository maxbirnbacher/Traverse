package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"traverse/pkg/utils"
	"traverse/pkg/logging"
	"traverse/pkg/proxy"
)

var redirects = make(map[string][]string)

var workers = []string{}

var token = utils.GenerateUUID()

func main() {
	// Redirect url
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requesting ", r.RequestURI, " from ", r.RemoteAddr)
		requested_path := r.URL.Path[1:]           // Get the requested path from the path
		paths, exists := redirects[requested_path] // Check if the path exists in the redirects map
		if !exists || len(paths) == 0 {
			http.NotFound(w, r)
			fmt.Println("Path not found")
			return
		}

		selectedUrl := paths[rand.Intn(len(paths))] // Select a random URL from the list of URLs
		fmt.Println("Redirecting request to ", selectedUrl)
		http.Redirect(w, r, selectedUrl, http.StatusFound)
	})

	// welcome url
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		//log the request
		fmt.Println("Requesting /welcome from ", r.RemoteAddr)
		//return welcome message
		fmt.Fprintf(w, "Welcome to Traverse!\n")
		//create a new UUID from the utils package
		uuid := utils.GenerateUUID()
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

		// check if the security token is correct
		if r.Form.Get("token") != token {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		//generate a fake URI path
		path, url := utils.GeneratePathAndURL(r)
		//add the URL to the redirects map
		redirects[path] = append(redirects[path], url)
		//return the URL
		fmt.Println("Generated URL: /", path, " to ", url)
		fmt.Fprintf(w, "Generated URL: /%s to %s\n", path, url)
	})

	// get the redirect url
	http.HandleFunc("/redirections", func(w http.ResponseWriter, r *http.Request) {
		//log the request
		fmt.Println("Requesting /redirections from ", r.RemoteAddr)
		//return the redirect map to the worker as a JSON response
		response := "{"
		for key, value := range redirects {
			response += fmt.Sprintf("\"%s\": %s, ", key, value)
		}
		response = response[:len(response)-2] + "}"
		fmt.Fprintf(w, response)
	})

	// Start the HTTP server
	output := "Server is running on port 8080"
	fmt.Println(output)
	fmt.Println("Use the security token to create a initial url: ", token)
	log.Fatal(http.ListenAndServe(":8080", nil))
}




