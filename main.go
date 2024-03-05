package main

import (
    "fmt"
    "os"
    "text/template"
    "net/http"
)

type Note struct{
    Title string
    Content string
}

type User struct{
    notes []Notes
    id string
    name string
}

func main(){
    note Note{} 
    mux := http.NewServeMux()

    mux.HandleFunc("POST /notes", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("Notes Created"))
    })

    
    mux.HandleFunc("GET /note/{id}", func(w http.ResponseWriter, r *http.Request){
        note_id := r.PathValue("id")
        fmt.Fprintf(w, "Note Id: %s", note_id)
    })
    
    fmt.Printf("Listening To: port:6969")
    http.ListenAndServe(":6969", mux)
}

