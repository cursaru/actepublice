package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	const numarport string = ":7000"
	//contor := 0


	r := mux.NewRouter()

	fmt.Printf("Starting...")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	x, err := fmt.Fprintf(w, "ActePublice.Ro %d", contor)

	// 	if err != nil {

	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(x)
	// 	}
	// 	contor ++
	// })

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request){
		
		parametri := mux.Vars(r)

		fmt.Println(parametri["title"])

		x, err := fmt.Fprintf(w, parametri["title"])

		if err != nil{
			fmt.Println(err)
		} else {
			fmt.Println(x)
		}

	})

	r.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request){

		x, err := fmt.Fprintf(w, "pagina main")

		if err != nil{
			fmt.Println(err)
		} else {
			fmt.Println(x)
		}

	})
	




	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	_ = http.ListenAndServe(numarport, r)  // daca se merge pe cale obisnuita aici nu mai este nil ci r


}
