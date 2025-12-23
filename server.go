package httplusplus

import (
	"fmt"
	"log"
	"net/http"
)

func StartApp(address string, app *AppRouter) {
	fmt.Printf("Server running at http://%s\n", address)
	if err := http.ListenAndServe(fmt.Sprintf("%s", address), app.Mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
