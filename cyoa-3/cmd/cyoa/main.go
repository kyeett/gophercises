package main

import (
	"net/http"
	"os"

	cyoa "github.com/kyeett/gophercises/cyoa-3"
)

func main() {

	filename := "gopher.json"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	handler := cyoa.NewHandler(story)

	//	tmpl.Execute(os.Stdout, story["intro"])
	http.ListenAndServe(":8080", handler)
}
