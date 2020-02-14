package main

import (
	"flag"
	"fmt"
	//"time"

	"github.com/nunosusanoribeiro/HTTPcompare/client"
	"github.com/nunosusanoribeiro/HTTPcompare/server"
	//"github.com/nunosusanoribeiro/HTTPcompare/internal/structs"

)

// Instantiates the Config struct to myConfig
//var myConfig structs.Config

func main() {

	// Populates the Config struct and parses CLI params
/*	
	myConfig.Verbose = flag.Bool("v", false, "verbose")
	myConfig.HttpVersion = flag.String("http", "1.1", "1.1, 2, 3")
	myConfig.Host = flag.String("host", "", "To connect to {msa,msb,msc,msd}.jpbourbon-httpspeed.io")
	myConfig.ServerName = flag.String("nameserv", "", "Server name for certificates {msa,msb,msc,msd}.jpbourbon-httpspeed.io")
	myConfig.IPServer = flag.String("ipserv", "", "IP address of server")
	myConfig.DBAddress = flag.String("dbhost", "127.0.0.1", "DB host address (only M-A)")
	myConfig.DBPort = flag.Int("dbport", 3306, "DB port (only M-A)")
	myConfig.DBUsername = flag.String("dbuser", "root", "DB username (only M-A)")
	myConfig.DBPassword = flag.String("dbpass", "", "DB password (only for M-A)")
	myConfig.DBName = flag.String("dbname", "httpcompare", "DB name (only for M-A)")
	myConfig.Delay = flag.String("delay", "0", "Request delay in miliseconds")

	flag.Parse()
*/

	
	// Start processes
	runHttp()
}

func runHttp() {
	
	var HttpVersion = flag.String("http", "1.1", "HTTP version - 1.1, 2, 3")
	var IPServer = flag.String("ipserv", "localhost", "IP address of server")
	var Port = flag.String("port", "8443", "Port of server")
	var TLSversion = flag.String("tls", "1.3", "TLS version - 1.2, 1.3")

	flag.Parse()

	fmt.Printf("%s\n", *HttpVersion)
	fmt.Printf("%s\n", *IPServer)
	fmt.Printf("%s\n", *Port)
	fmt.Printf("%s\n", *TLSversion)


	go server.StartServer(*HttpVersion, *IPServer, *Port, *TLSversion)

	for i := 0; i < 100 ; i++ {

		client.StartClient(*HttpVersion, *IPServer, *Port, *TLSversion)

	}
	
	

}