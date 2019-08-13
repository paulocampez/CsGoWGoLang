package main


import (
	"net/http"
)

func main(){
//https://github.com/go-chi/chi
	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}