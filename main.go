
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ai-agent-platform/handlers"
)

// Application entry point 
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	mux.HandleFunc("POST /api/agents", handlers.CreateAgent)
	mux.HandleFunc("GET /api/agents", handlers.ListAgents)
	mux.HandleFunc("POST /api/agents/{id}/runs", handlers.CreateRun)
	mux.HandleFunc("GET /api/runs/{id}", handlers.GetRun)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Println("AI Agent API running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}