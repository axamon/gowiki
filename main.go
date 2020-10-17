package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var address = flag.String("addr", ":8080", "Server address")
	flag.Parse()

	r := http.NewServeMux()

	fsmedia := http.StripPrefix("/media/", http.FileServer(http.Dir("./media")))
	r.Handle("/media/", fsmedia)

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/view/", makeHandler(viewHandler))
	r.HandleFunc("/edit/", makeHandler(editHandler))
	r.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(*address, r))
}
