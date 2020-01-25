package main


import ( 
	"fmt"
	"net/http"
	"html/template"
)



func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles("views/Index.html"))
	data:=""
	var err = tmp.ExecuteTemplate(w,"Index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func HandlerGaleri(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles("views/Galeri.html"))
	data:=""
	var err = tmp.ExecuteTemplate(w,"Galeri", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func main() {

    //file static
    http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("assets"))))

    //router
	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/galeri", HandlerGaleri)

	fmt.Println("starting web server at http://localhost:8000/")
    http.ListenAndServe(":8000", nil)

}