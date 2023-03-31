package main

// import properties
import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverlist = []*server{
		newServer("http://127.0.0.1:5000/"),
		newServer("http://127.0.0.1:5001/"),
		newServer("http://127.0.0.1:5002/"),
		newServer("http://127.0.0.1:5003/"),
		newServer("http://127.0.0.1:5004/"),
	}
	lastServerIndex = 0
)

func main() {
	http.HandleFunc("/", forward)
	//Errorformats
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forward(r http.ResponseWriter, w *http.Request) {
	server := getServer()
	server.ReverseProxy.ServeHTTP(r, w)
}

func getServer() *server {
	nextIndex := (lastServerIndex + 1) % len(serverList)
	server := serverList[nextIndex]
	lastServerIndex = nextIndex
	return server
}

func getHealthyServer() *server {
	for i := 0; i < len(serverList); i++ {
		server := getServer()
		if server.Health {
			return server, nil
		}
	}
}

func createHost(urlStr string) *httputil.ReverseProxy {
	u, _ := url.Parse(urlStr)
	return httputil.NewSingleHostReverseProxy(u)
}
