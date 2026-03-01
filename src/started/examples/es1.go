package examples

import (
	"fmt"
	"net/http"
)

func FirstEx () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1 style='color: red; font-size: 5rem'> Hello World! </h1>")
	})

	http.ListenAndServe(":8080", nil)
}