package main
import 
(
  "html/template"
  "log"
  "net/http"
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
   fileServer := http.FileServer(http.Dir("static"))
   http.Handle("/static/", http.StripPrefix("/static/", fileServer))
   http.HandleFunc("/", renderTemplate)
   err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
   if err != nil {
         log.Fatal("error starting http server : ", err)
         return
   }
}