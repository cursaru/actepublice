package renders

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string, data interface{}) {

	var tmpl *template.Template
	var err error

	//verificam daca exista template-ul in tc
	// inMap este un bool
	_, inMap := tc[t]

	if !inMap {
		// trebuie facut ceva daca nu exista in cache, adica este false

		log.Println("Creating template and adding to cache...")
		err = createTemplateCache(t)

		if err != nil {
			log.Println(err)
		}

	} else {
		// trebuie facut ceva daca exista in cache, adica este true

		log.Println("serving ", t, " from cache")

		// ca sa vad eu cum e treba si ce e in cache, dar nu e obligatoriu
		for element, _ := range tc {

			log.Println(element)
		}
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {

	templates := []string{

		fmt.Sprintf("./pagini/%s", t),
		"./pagini/base.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {

		return err
	}

	tc[t] = tmpl

	return nil
}
