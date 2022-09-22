package api

import (
	"net/http"
)

func (s *Server) handleGreeter() http.HandlerFunc {
	// return func(w http.ResponseWriter, r *http.Request) {
	// 	b, err := io.ReadAll(io.LimitReader(r.Body, 100000))
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	name := string(b)
	// 	w.WriteHeader(http.StatusOK)
	// 	fmt.Fprintf(w, "Hello there, %s", name)
	// }
	return func(w http.ResponseWriter, r *http.Request) {
		greetings := []struct {
			Language string
			Greeting string
		}{
			{"English", "Hello"},
			{"Arabic", "Salem"},
		}
		s.respond(w, r, greetings, http.StatusOK)
	}
}
