package examples

import (
	"fmt"
	"net/http"
	"strconv"
)

func ThirdEx () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			count, _ := strconv.Atoi(r.FormValue("counter"))
			count++

			fmt.Fprintf(w, `
				<body>
					<form action="/" method="POST">
						<label>Counter</label>
						<input name="counter" value="%d" readonly>
						<button type="submit">Incrementa</button>
					</form>
					<a href="/">Reset</a>
				</body>
			`, count)

			return
		}

		fmt.Fprintf(w, `
			<body>
				<form action="/" method="POST">
					<label>Counter</label>
					<input name="counter" value="0" readonly>
					<button type="submit">Incrementa</button>
				</form>
			</body>
			<a href="/">Reset</a>
		`)
	})

	http.ListenAndServe(":8080", nil)
}