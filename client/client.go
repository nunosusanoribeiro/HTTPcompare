package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func StartClient(httpV string, ip string, port string, tlsV string) {
	// Request /hello over port 8080 via the GET method
	//r, err := http.Get("http://localhost:8080/hello")

	caCert, err := ioutil.ReadFile("internal/certs/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	var conn = "https://"+ip+":"+port+"/hello"
	// Request /hello over HTTPS port 8443 via the GET method
	r, err := client.Get(conn)

	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}
