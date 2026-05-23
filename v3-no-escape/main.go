package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	fsys, err := NewCustomFS(".", true, false)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveCustomFS(fsys, w, r)
	})
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Custom File System struct, methods and associated functions
type customFileSystem struct {
	innerFS       fs.FS
	dotFileAccess bool
}

func NewCustomFS(rootPath string, rootEscape bool, dotFileAccess bool) (*customFileSystem, error) {
	var innerFS fs.FS
	if rootEscape {
		innerFS = os.DirFS(rootPath)
	} else {
		root, err := os.OpenRoot(rootPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open restricted root: %w", err)
		}
		innerFS = root.FS()
	}

	return &customFileSystem{
			innerFS:       innerFS,
			dotFileAccess: dotFileAccess},
		nil
}

func serveCustomFS(fsys fs.FS, w http.ResponseWriter, r *http.Request) {
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
		urlPath = path.Join(urlPath, "index.html")
		if _, err := fs.Stat(fsys, urlPath); err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
	}
	http.ServeFileFS(w, r, fsys, urlPath)
}

func (fsys *customFileSystem) Open(name string) (fs.File, error) {
	if !fsys.dotFileAccess && containsDotFile(name) {
		return nil, fs.ErrPermission
	}
	return fsys.innerFS.Open(name)
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
