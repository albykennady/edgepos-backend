package httpapi

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO 1: Set the Content-Type header to application/json.

	w.Header().Set("Content-Type", "application/json")

	// TODO 2: Set the HTTP status code to 200 OK.
	w.WriteHeader(http.StatusOK)

	// TODO 3: Write the JSON response body {"status":"ok"}.

	w.Write([]byte(`{"status":"ok"}`))

}
