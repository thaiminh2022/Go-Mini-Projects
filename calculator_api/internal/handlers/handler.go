package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thaiminh2022/calculator_api/api"
)

const CONTENT_TYPE string = "Content-TYPE"
const JSON_MODE string = "application/json"

func Handler(mux *http.ServeMux) {
	mux.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		input := api.Input{}
		json.NewDecoder(r.Body).Decode(&input)

		answer := api.Answer{Answer: input.A + input.B}
		w.Header().Set(CONTENT_TYPE, JSON_MODE)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&answer)
	})
	mux.HandleFunc("POST /minus", func(w http.ResponseWriter, r *http.Request) {
		input := api.Input{}
		json.NewDecoder(r.Body).Decode(&input)

		answer := api.Answer{Answer: input.A - input.B}
		w.Header().Set(CONTENT_TYPE, JSON_MODE)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&answer)
	})
	mux.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		input := api.Input{}
		json.NewDecoder(r.Body).Decode(&input)

		answer := api.Answer{Answer: input.A * input.B}
		w.Header().Set(CONTENT_TYPE, JSON_MODE)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&answer)
	})
	mux.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		input := api.Input{}
		json.NewDecoder(r.Body).Decode(&input)
			w.Header().Set(CONTENT_TYPE, JSON_MODE)
		switch input.B {
		case 0:
            err := api.Error{Message: "Cannot divide by 0 (b value must be non 0)"}
			w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(&err)
		default:
			answer := api.Answer{Answer: input.A / input.B}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&answer)
		}

	})
}
