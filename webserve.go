package main

import "flag"
import "fmt"
import "log"
import "net/http"

func main() {
	var ipStringPtr = flag.String("ip", "0.0.0.0", "IP address to bind to")
	var portStringPtr = flag.String("port", "8080", "port number to bind to")
	flag.Parse()

	var bindStr = fmt.Sprintf("%s:%s", *ipStringPtr, *portStringPtr);

	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println(fmt.Sprintf("Serving current directory on http://%s/", bindStr))
	log.Fatal(http.ListenAndServe(bindStr, nil))
}
