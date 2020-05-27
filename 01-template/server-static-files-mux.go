package main
import 
(
  "html/template"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)
const
(
  CONN_HOST = "localhost"
  CONN_PORT = "8080"
)
type Person struct {
    Name string
    Id string
    Age string
}
func renderTemplate(w http.ResponseWriter, r *http.Request) {
  person := Person{Name: "Jocelyn", Id: "5", Age : "22" }
  parsedTemplate, _ := template.ParseFiles("templates/02-template.html")
  err := parsedTemplate.Execute(w, person)
  if err != nil {
     log.Printf("Error occured  while excuting the template : ", err)
     return
  }
}
func main() {
   router := mux.NewRouter()
   router.HandleFunc("/", renderTemplate).Methods("GET")
   router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
   err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
   if err != nil {
         log.Fatal("error starting http server : ", err)
         return
   }
}