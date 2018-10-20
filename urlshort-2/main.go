package main

import (
	"fmt"
	"log"
	"net/http"
)

func newDefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	return mux
}

func main() {

	mux := newDefaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
  `
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "No such path :-)")
}
