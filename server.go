package jbls

import (
	"net/http"
)

func readme_handler()  {
	
}

func server()  {
	http.HandleFunc("/", readme_handler)
}

func main() {

}
