package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//error
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// creating a server instance
	mux := http.NewServeMux()
	//serving static files
	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	//managing req
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	//handting sockets
	addr := flag.String("addr", ":4000", "the server socket address")
	flag.Parse()
	infoLog.Printf("starting the server at %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
