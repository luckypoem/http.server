package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	flListen  = flag.String("l", ":8080", "listen address")
	flVerbose = flag.Bool("v", false, "be verbose")
)

func mdLogs(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.RequestURI())
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root, _ = os.Getwd()
	} else {
		root, _ = filepath.Abs(root)
	}
	log.Println("Root", root)

	handler := http.FileServer(http.Dir(root))
	if *flVerbose {
		handler = mdLogs(handler)
	}
	http.Handle("/", handler)
	log.Println("Listen and serve on", *flListen)
	if err := http.ListenAndServe(*flListen, nil); err != nil {
		log.Fatalln(err)
	}
}
