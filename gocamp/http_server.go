package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverApp() {
	app := http.NewServeMux()
	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello world, app")
	} )
	fmt.Println("listening 8000...")
	if err := http.ListenAndServe(":8000", app); err != nil {
		log.Fatal(err)
	}
}

func serverDebug() {
	app := http.NewServeMux()
	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello world, debug")
	} )
	fmt.Println("listening 8001...")
	if err := http.ListenAndServe(":8001", app); err != nil {
		log.Fatal(err)
	}
}
func main(){
	go serverApp()
	go serverDebug()

	select {}
}