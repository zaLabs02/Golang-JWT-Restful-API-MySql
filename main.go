package main

import (
	"fmt"
	"log"
	"login-register/config"
	"login-register/routes"
	"login-register/seeds"
	"net/http"
)

func main() {
	r := routes.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 8080...")
	seeds.Load(config.Database)
	log.Fatal(http.ListenAndServe(":8080", r))
}
