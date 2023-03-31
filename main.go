package main

//import packages
import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// serverList
var serverList = []string{
	"http://127.0.0.1:5000/",
	"http://127.0.0.1:5001/",
	"http://127.0.0.1:5002/",
	"http://127.0.0.1:5003/",
	"http://127.0.0.1:5004/",
}

// lastserverIndex
var lastServerIndex = 0

func main() {
	http.HandleFunc("/", forward)
	if http.ListenAndServe(":8000", nil) != nil {
		log.Fatal("server error")
	}
}

func forward(w http.ResponseWriter, r *http.Request) {
	urls := getURl()
	reverseProxy := httputil.NewSingleHostReverseProxy(urls)
	log.Printf("Routing request switch to the url:%s", urls.String())
	reverseProxy.ServeHTTP(w, r)
}

func getURl() *url.URL {
	url, _ := url.Parse(serverList[(lastServerIndex+1)%len(serverList)])
	lastServerIndex += 1
	if lastServerIndex > 5 {
		lastServerIndex = 0
	}
	return url
}
