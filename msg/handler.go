package msg

import (
	"net/http"
	"log"
)

func Msg(svcName string) []byte {
	return []byte("hello from " + svcName + "\n")
}

func MsgHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			greeting := Msg("service A")
			w.Write(greeting)
		},
	)
}

func ConditionalHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			response := Msg("service A")

			// Check to see if the request has an arbitrary URL param set
			// A request to http://localhost:8080?key=test will skip this block
			keys, ok := r.URL.Query()["key"]
			if !ok || len(keys[0]) < 1 {
				log.Println("Url Param 'key' is missing")
				response = Msg("unknown service")
			}

			w.Write(response)
		},
	)
}