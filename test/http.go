package main

import (
	"log"
	"net/http"
	"time"
)

// v1

//func main (){
//   http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
//   	 w.Write([]byte("Hllo,this is version 1!"));
//   });
//
//   http.HandleFunc("/bye",sayBye);
//   log.Println("starting server ... v1");
//   log.Fatal(http.ListenAndServe(":4000",nil));
//}
//
//func sayBye(w http.ResponseWriter,r *http.Request){
//	w.Write([]byte("BYE!"));
//}

//v2

type myHandler struct {
}

func main() {
	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	server.Handler = mux

	http.HandleFunc("/bye", sayBye)
	log.Println("starting server ... v2")
	//log.Fatal(http.ListenAndServe(":4000", mux))
	log.Fatal(server.ListenAndServe())
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hllo,this is version 2!" + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BYE 2!"))
}
