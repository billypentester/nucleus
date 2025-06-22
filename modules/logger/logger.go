package logger

import (
	"log"
	"net/http"
)

func Log(r *http.Request) {

	// TODO : Currently it only logs the method and path. It should be extended more.
	log.Printf("%s %s\n", r.Method, r.URL.Path)

}
