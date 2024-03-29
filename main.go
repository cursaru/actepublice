package main

import (
	"actepublice/renders"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	const numarport string = ":7000"
	//contor := 0

	r := mux.NewRouter()

	fmt.Printf("Starting...")

	r.HandleFunc("/articole/{articol}", func(w http.ResponseWriter, r *http.Request) {

		parametri := mux.Vars(r)

		//fmt.Println(parametri["articol"])

		data := parametri // structures.EmptyInteger

		//x, err := fmt.Fprintf(w, parametri["articol"])

		renders.RenderTemplate(w, "post.html", data)

		//if err != nil{
		//	fmt.Println(err)
		//} else {
		//	fmt.Println(x)
		//}

	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		x, err := fmt.Fprintf(w, "pagina main")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(x)
		}

	})

	//fs := http.FileServer(http.Dir("static/"))

	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	_ = http.ListenAndServe(numarport, r) // daca se merge pe cale obisnuita aici nu mai este nil ci r

}
