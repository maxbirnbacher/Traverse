package proxy

import (
    "net/http"
    "net/http/httputil"
    "net/url"
)

func ProxyRequest(w http.ResponseWriter, r *http.Request, protocol string, ip string, port string) {
    // Parse the destination server's URL
    target, _ := url.Parse(protocol + "://" + ip + ":" + port)

    // Create a new reverse proxy that forwards requests to the destination server
    proxy := httputil.NewSingleHostReverseProxy(target)

    // Update the headers to allow for SSL redirection
    r.URL.Host = target.Host
    r.URL.Scheme = target.Scheme
    r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
    r.Host = target.Host

    // Use the reverse proxy to serve the request
    proxy.ServeHTTP(w, r)
}