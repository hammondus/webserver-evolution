package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServerFS(os.DirFS(".")) // introduced in go1.22
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
