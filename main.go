package main

import (
	"fmt"
	"net/http"
)

func main() {

	const numarport string = ":7000"
	contor := 0

	fmt.Printf("Starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		x, err := fmt.Fprintf(w, "ActePublice.Ro %d", contor)

		if err != nil {

			fmt.Println(err)
		} else {
			fmt.Println(x)
		}
		contor ++
	})

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	_ = http.ListenAndServe(numarport, nil)

}
