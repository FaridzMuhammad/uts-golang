package main

import (
	"fmt" 
	"mahasiswa" 
	"net/http"
)

func main() {
	http.HandleFunc("/mahasiswa/data", mahasiswa.GetMahasiswaHandler)
	http.HandleFunc("/mahasiswa/find", mahasiswa.FindMahasiswaHandler)
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}