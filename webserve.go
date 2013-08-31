package main

import "flag"
import "fmt"
import "log"
import "net/http"

func main() {
	var specStringPtr = flag.String("addr", ":8080", "[address][:port] to serve from")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println(fmt.Sprintf("Serving current directory on %s", *specStringPtr))
	log.Fatal(http.ListenAndServe(*specStringPtr, nil))
}
