package jbls

import (
	"log"
	"flag"
	"os"
)

func main() {
	ip := flag.String("ip", "0.0.0.0", "Bind your ip address.")
	port := flag.String("port", "8000", "Bind your port.")
	keyPath := flag.String("key", "", "Private key for the license server.")
	flag.Parse()

	log.Println("Bind ip address to: " + *ip)
	log.Println("Bind port to: " + *port)
	if _, err := os.Stat(*keyPath); os.IsNotExist(err) {
		log.Printf("Key file(%s) is not exist.", *keyPath)
		return
	}
}
