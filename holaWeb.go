package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "!Holaaaaaaaa")
	})
	http.HandleFunc("/nombre", miNombre)

	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println()
	io.WriteString(w, "Hola mundo")
}

func miNombre(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Mi nombre es WDavid Calsin")
}
