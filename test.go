package main

import (
	"fmt"
	"html"
	"net/http"
)

func req() {
	// make a sample server

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(":8080", nil)
	//hitting this api with my phone entering pc's ip and port will not work if the eBPF program is running
	//because the eBPF program will drop the packet

}
