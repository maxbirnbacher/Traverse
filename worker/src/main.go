package main

import (
	"io"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the original request details
		originalRequestURL := r.Header.Get("X-Original-Request-URL")
		originalRequestMethod := r.Header.Get("X-Original-Request-Method")
		
		// Create a new request to the hidden URL
		req, err := http.NewRequest(originalRequestMethod, originalRequestURL, r.Body)
		if err != nil {
			http.Error(w, "Failed to create new request", http.StatusInternalServerError)
			return
		}
		
		// Copy headers from the original request to the new request
		for headerName, headerValues := range r.Header {
			for _, headerValue := range headerValues {
				req.Header.Set(headerName, headerValue)
			}
		}
		
		// Send the request to the hidden URL
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Failed to forward request", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()
		// Write the response from the hidden URL back to the master server
		copyHeader(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})
	
	http.ListenAndServe(":8081", nil)
}

func copyHeader(dst, src http.Header) {
    for k,
    vv := range src {
        for _,
        v := range vv {
            dst.Add(k, v)
        }
    }
}
