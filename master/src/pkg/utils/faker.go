package utils

import (
    "fmt"
    "net/http"
    "strings"
)

func GeneratePathAndURL(r *http.Request) (string, string) {
    // TODO: NEEDS REWRITE
    
    // // Generate a fake URI path
    // path := faker.Word()
    // // Use only the part after the last slash
    // path = path[strings.LastIndex(path, "/")+1:]
    // fmt.Println("Generated path: ", path)
    // // Get the URL from the query parameters
    // url := r.Form.Get("url")

    // return path, url
}