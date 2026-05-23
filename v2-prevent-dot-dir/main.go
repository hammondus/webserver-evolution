package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	fsys := customFileSystem{
		FS:            os.DirFS("."),
		dotFileAccess: false,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveCustomFS(&fsys, w, r)
	})
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Custom File System struct, methods and associated functions
type customFileSystem struct {
	fs.FS
	dotFileAccess bool
}

func serveCustomFS(fsys *customFileSystem, w http.ResponseWriter, r *http.Request) {
	urlPath := path.Clean(r.URL.Path)
	urlPath = strings.TrimPrefix(urlPath, "/")

	// Prevent an empty path
	if urlPath == "" {
		urlPath = "."
	}

	stat, err := fs.Stat(fsys, urlPath)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if stat.IsDir() {
		indexPath := path.Join(urlPath, "index.html")
		if _, err := fs.Stat(fsys, indexPath); err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.ServeFileFS(w, r, fsys, indexPath)
		return
	}
	http.ServeFileFS(w, r, fsys, urlPath)
}

func (fsys *customFileSystem) Open(name string) (fs.File, error) {
	if !fsys.dotFileAccess && containsDotFile(name) {
		return nil, fs.ErrPermission
	}
	return fsys.FS.Open(name)
}

func containsDotFile(name string) bool {
	parts := strings.SplitSeq(name, "/")
	for part := range parts {
		if part != "." && strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}
