package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It works!"))
}

func pingHandler(keyData []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		salt := r.URL.Query().Get("salt")
		data := "<PingResponse><message></message><responseCode>OK</responseCode><salt>" +
			salt +
			"</salt></PingResponse>"
		signature := sign(data, keyData)
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		w.Write([]byte("<!-- " + signature + " -->\n" + data))

		log.Print("Pong!")
	}
}

func prolongTicketHandler(keyData []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		salt := r.URL.Query().Get("salt")
		data := "<ProlongTicketResponse><message></message><responseCode>OK</responseCode><salt>" +
			salt +
			"</salt><ticketId>1</ticketId></ProlongTicketResponse>"
		signature := sign(data, keyData)
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		w.Write([]byte("<!-- " + signature + " -->\n" + data))

		log.Print("Prolong ticket!")
	}
}

func releaseTicketHandler(keyData []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		salt := r.URL.Query().Get("salt")
		data := "<ReleaseTicketResponse><message></message><responseCode>OK</responseCode><salt>" +
			salt +
			"</salt></ReleaseTicketResponse>"
		signature := sign(data, keyData)
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		w.Write([]byte("<!-- " + signature + " -->\n" + data))

		log.Print("Release ticket!")
	}
}

func obtainTicketHandler(keyData []byte, userName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		salt := r.URL.Query().Get("salt")
		name := r.URL.Query().Get("userName")
		if userName != "" {
			name = userName
		}
		const prolongationPeriod = "607875500"
		data := "<ObtainTicketResponse><message></message><prolongationPeriod>" +
			prolongationPeriod +
			"</prolongationPeriod><responseCode>OK</responseCode><salt>" +
			salt +
			"</salt><ticketId>1</ticketId><ticketProperties>licensee=" +
			name +
			"\tlicenseType=0\t</ticketProperties></ObtainTicketResponse>"
		signature := sign(data, keyData)
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		w.Write([]byte("<!-- " + signature + " -->\n" + data))

		log.Print("Obtain ticket!")
	}
}

func server(host string, port string, keyData []byte, name string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.Handle("/rpc/ping.action", pingHandler(keyData))
	mux.Handle("/rpc/prolongTicket.action", prolongTicketHandler(keyData))
	mux.Handle("/rpc/releaseTicket.action", releaseTicketHandler(keyData))
	mux.Handle("/rpc/obtainTicket.action", obtainTicketHandler(keyData, name))

	log.Println("Starting server!")
	http.ListenAndServe(host+":"+port, mux)
}
