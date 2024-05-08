package main

import (
	"fmt"
	"os"
	"log"
	"math/rand"
	"net/http"
	"traverse/pkg/utils"
	"traverse/pkg/logging"
	"traverse/pkg/proxy"
)

var workers = []string{}

var token = utils.GenerateUUID()

func main() {
	// Get the IP and port for the HorizonBackbone server from the environment variables
	hb_ip := os.Getenv("HB_IP")
	hb_c2_port := os.Getenv("HB_C2_PORT")
	hb_log_port := os.Getenv("HB_LOG_PORT")
	port := os.Getenv("TRAVERSE_PORT")

	// Set up the logging
	logging.SetUpLogging(hb_ip, hb_log_port)

	// redirect the requests to the HorizonBackbone server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requesting ", r.RequestURI, " from ", r.RemoteAddr)
		logging.Info("Requesting " + r.RequestURI + " from " + r.RemoteAddr)
		//check if the request was made with http or https
		protocol := "http"
		if r.TLS != nil {
			protocol = "https"
		}
		logging.Info("Redirecting request to " + hb_ip + ":" + hb_c2_port)
		proxy.ProxyRequest(w, r, protocol, hb_ip, hb_c2_port)
	})

	// Start the HTTP server
	output := "Server is running on port "+port
	logging.Info(output)
	logging.Info("Use the security token to register worker: " + token)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}




