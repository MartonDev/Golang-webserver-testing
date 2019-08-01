package main

import (
	"net/http"
	"os"

	"github.com/martondev/Golang-webserver-testing/api"
)

type HTMLDir struct {
	d http.Dir
}

func main() {

	fs := http.FileServer(HTMLDir{http.Dir("public/")})

	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/api/", api.API)
	http.ListenAndServe(":8000", nil)

}

func (d HTMLDir) Open(name string) (http.File, error) {

	f, err := d.d.Open(name)

	if os.IsNotExist(err) {

		if f, err := d.d.Open(name + ".html"); err == nil {

			return f, nil

		}

	}

	return f, err

}
