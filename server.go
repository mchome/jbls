package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It works!"))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Pong!")
	w.Write([]byte("Pong!"))
}

func releaseTicketHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Release it!")
}

func obtainTicketHandler(keyPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		salt := r.URL.Query().Get("salt")
		name := r.URL.Query().Get("userName")
		const prolongationPeriod = "607875500"
		data := "<ObtainTicketResponse><message></message><prolongationPeriod>" +
			prolongationPeriod +
			"</prolongationPeriod><responseCode>OK</responseCode><salt>" +
			salt +
			"</salt><ticketId>1</ticketId><ticketProperties>licensee=" +
			name +
			"\tlicenseType=0\t</ticketProperties></ObtainTicketResponse>"
		signature := obtainTicket(data, keyPath)
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		w.Write([]byte("<!-- " + signature + " -->\n" + data))
	}
}

func server(host string, port string, keyPath string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/rpc/ping.action", pingHandler)
	mux.HandleFunc("/rpc/releaseTicket.action", releaseTicketHandler)
	mux.Handle("/rpc/obtainTicket.action", obtainTicketHandler(keyPath))

	log.Println("Starting server!")
	http.ListenAndServe(host+":"+port, mux)
}
