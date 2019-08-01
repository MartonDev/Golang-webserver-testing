package api

import (
	"fmt"
	"net/http"
)

func API(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)

}
