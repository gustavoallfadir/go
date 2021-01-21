package main

import (
	"fmt"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	fmt.Println("Escuchando peticiones en el puerto 3000")
	http.ListenAndServe(":3000",mux)
}