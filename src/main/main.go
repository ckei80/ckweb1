package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type MainHandler struct {
}

func (this *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404 error - " + http.StatusText(404)))
	} else {
		w.Write(data)
	}

}

func main() {
	http.Handle("/", new(MainHandler))
	http.ListenAndServe(":8080", nil)
}
