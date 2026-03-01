package examples

import (
	"fmt"
	"net/http"
)

func SecondEx () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("q") != "" {
			fmt.Fprintf(w, `
				<body>
					<h1 align="center" style="color: red; font-size: 5rem"> Ciao pirla, %s </h1>
				</body>
			`, r.FormValue("q"))

			return
		}

		fmt.Fprintf(w, `
			<body>
				<form action="/" method="GET">
					<label>Come ti chiami? </label>
					<input name="q">
					<button type="submit">Invia</button>
				</form>
			</body>
		`)
	})

	http.ListenAndServe(":8080", nil)
}