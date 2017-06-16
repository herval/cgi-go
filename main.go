package main

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// load compressed resources from the "bundle"
func prodHandler(w http.ResponseWriter, r *http.Request) {
	resource := fmt.Sprintf("resources/%s", r.URL.Path[1:])
	measure(func() {
		if resource == "resources/" {
			resource = "resources/index.html" // default "/"
		}

		ext := filepath.Ext(resource)
		mimeType := mime.TypeByExtension(ext)
		bytes, err := Asset(resource)

		fmt.Println("Requested", resource, mimeType)

		if err != nil {
			fmt.Println(AssetNames())
			http.NotFound(w, r)
		} else {
			if mimeType != "" {
				w.Header().Set("Content-Type", mimeType)
			}
			w.Write(bytes)
		}
	})
}

// Handler for development mode - just use the resources under resources/
func devHandler(w http.ResponseWriter, r *http.Request) {
	resource := fmt.Sprintf("resources/%s", r.URL.Path[1:])
	measure(func() {
		fmt.Println("Requested", resource)
		http.ServeFile(w, r, resource)
	})
}

func measure(executor func()) {
	start := time.Now()
	executor()
	elapsed := time.Since(start)
	fmt.Printf("Served in %s\n\n", elapsed)
}

func main() {
	mode := "prod"
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}

	if mode == "dev" {
		fmt.Println("Using resources on resources/ folder")
		http.HandleFunc("/", devHandler)
	} else {
		fmt.Println("Using embedded resources")
		http.HandleFunc("/", prodHandler)
	}
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server running on port 8080")
}

func resourceExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
