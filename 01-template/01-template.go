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
type Person struct{
  Id string
  Name string
}
func renderTemplate(w http.ResponseWriter, r *http.Request) {
  person := Person{Id: "15", Name: "Aaron"}
  parsedTemplate, _ := template.ParseFiles("templates/01-template.html")
  err := parsedTemplate.Execute(w, person)
  if err != nil {
     log.Printf("Error occured while executing the template or writing its output : ", err)
     return
  }
}
func main(){
  http.HandleFunc("/", renderTemplate)
  err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
  if err != nil {
    log.Fatal("Error starting http server : ", err)
    return
  }
}