package httpapi

import "net/http"

func ReadyHandler(w http.ResponseWriter, r *http.Request) {

	// TODO 1: Set Content-Type to application/json.
	w.Header().Set("Content-Type", "application/json")

	// TODO 2: Return HTTP 200 OK.
	w.WriteHeader(http.StatusOK)
	// TODO 3: Return this JSON response:
	w.Write([]byte(`{"status":"ready"}`))
}
