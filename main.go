package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	host := flag.String("host", "127.0.0.1", "Bind your ip address.")
	port := flag.String("port", "8080", "Bind your port.")
	keyPath := flag.String("key", "", "Private key file path for the license server.")
	name := flag.String("name", "", "Give a fixed name to user.")
	flag.Parse()

	log.Printf("Bind to: %s:%s.", *host, *port)
	if _, err := os.Stat(*keyPath); os.IsNotExist(err) {
		log.Fatalf("Private Key file(%s) is not exist!", *keyPath)
	}
	isKey, keyData := validateKey(*keyPath)
	if (!isKey) {
		log.Fatal("Error private key!")
	}
	server(*host, *port, keyData, *name)
}
