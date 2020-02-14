package server

import (
	"io"
	"net/http"

	//"github.com/nunosusanoribeiro/HTTPcompare/client"
	//"github.com/nunosusanoribeiro/HTTPcompare/server"
	//"github.com/nunosusanoribeiro/HTTPcompare/internal/structs"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func StartServer(httpV string, ip string, port string, tls string) {
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	var conn = ip+":"+port

	// Listen to port 8080 and wait
	http.ListenAndServeTLS(conn,"C:/Users/Nuno/go/src/github.com/nunosusanoribeiro/HTTPcompare/internal/certs/cert.pem",
		"C:/Users/Nuno/go/src/github.com/nunosusanoribeiro/HTTPcompare/internal/certs/key.pem", nil)
}


