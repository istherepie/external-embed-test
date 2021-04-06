package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/istherepie/external-embed-test/pkg/ui"
)

func main() {
	// embedded http.Filesystem
	assets := ui.StaticAssets()

	// FS handler
	fs := http.FileServer(assets)

	// Start service
	host := "localhost"
	port := 8080

	address := fmt.Sprintf("%v:%d", host, port)

	log.Printf("Starting webserver on %v\n", address)
	serveErr := http.ListenAndServe(address, fs)

	if serveErr != nil {
		log.Fatal(serveErr)
	}

}
