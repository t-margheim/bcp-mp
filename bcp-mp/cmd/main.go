package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dimfeld/httptreemux"
)

func main() {
	router := httptreemux.NewContextMux()

	reactHandler := http.FileServer(reactFileSystem{http.Dir("./build")})
	router.GET("/*path", reactHandler.ServeHTTP)
	router.GET("/", reactHandler.ServeHTTP)

	portNumber := ":3000"
	if os.Getenv("PORT") != "" {
		portNumber = os.Getenv("PORT")
	}

	// Suggested to use a custom server with timeouts specified:
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	srv := &http.Server{
		Addr:         portNumber,
		Handler:      router,
		ReadTimeout:  60 * time.Second, // Default is unlimited
		WriteTimeout: 60 * time.Second, // Default is unlimited
	}
	log.Println("Service running...")
	log.Panic(srv.ListenAndServe())
}

type reactFileSystem struct {
	http.Dir
}

func (r reactFileSystem) Open(name string) (http.File, error) {
	// Point everything that isn't a file at index.html so React can determine it
	if name != "/" && !strings.Contains(name, ".") {
		name = "/index.html"
	}
	return r.Dir.Open(name)
}
