package main

import (

  "net/http"
  "os"
  "fmt"

)

type HTMLDir struct {

  d http.Dir

}

func main() {

  fs := http.FileServer(HTMLDir{http.Dir("public/")})

  http.Handle("/", http.StripPrefix("/", fs))
  http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {

    switch r.Method {

    case "GET":
      fmt.Fprintf(w, "GET")

    case "POST":
      fmt.Fprintf(w, "POST")

    default:
      fmt.Fprintf(w, "Other method")

    }

  })
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
