
package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"ai-agent-platform/models"
)

var (
	mu     sync.RWMutex
	agents = make(map[string]models.Agent)
	runs   = make(map[string]models.Run)
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func CreateAgent(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
		return
	}

	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Agent name is required",
		})
		return
	}

	agent := models.Agent{
		ID:          time.Now().Format("20060102150405.000000000"),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
	}

	mu.Lock()
	agents[agent.ID] = agent
	mu.Unlock()

	writeJSON(w, http.StatusCreated, agent)
}

func ListAgents(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	result := make([]models.Agent, 0, len(agents))
	for _, agent := range agents {
		result = append(result, agent)
	}
	mu.RUnlock()

	writeJSON(w, http.StatusOK, result)
}

func CreateRun(w http.ResponseWriter, r *http.Request) {
	agentID := r.PathValue("id")

	var input struct {
		Input string `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
		return
	}

	mu.Lock()
	agent, exists := agents[agentID]
	if !exists {
		mu.Unlock()
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Agent not found",
		})
		return
	}

	run := models.Run{
		ID:        time.Now().Format("20060102150405.000000000"),
		AgentID:   agent.ID,
		Status:    "queued",
		Input:     input.Input,
		CreatedAt: time.Now(),
	}

	runs[run.ID] = run
	mu.Unlock()

	writeJSON(w, http.StatusAccepted, run)
}

func GetRun(w http.ResponseWriter, r *http.Request) {
	runID := r.PathValue("id")

	mu.RLock()
	run, exists := runs[runID]
	mu.RUnlock()

	if !exists {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Run not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, run)
}