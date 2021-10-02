package main

import (
    "html/template"
    "log"
    "net/http"
    "time"
    "net"
)

type PageData struct{
    Title string
    Time string
    IPLocal string
}

func GetServerIP() string{
    conexao, erro := net.Dial("udp", "8.8.8.8:80")
    if erro != nil{
        return "Erro"
    }
    ip := conexao.LocalAddr().String()
    conexao.Close()
    return ip
}

func Handler(writer http.ResponseWriter, request *http.Request) {
    currentTime := time.Now()
    pageData := PageData{Title:"Super pagina"}
    pageData.Time = currentTime.String()
    pageData.IPLocal = GetServerIP()
    err := templates.ExecuteTemplate(writer, "page.html", pageData)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}
var templates = template.Must(template.ParseFiles("templates/page.html"))
func main() {
    http.HandleFunc("/", Handler)
    log.Fatal(http.ListenAndServe(":80", nil))
}
