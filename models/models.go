
package models

import "time"

type Agent struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Run struct {
	ID        string    `json:"id"`
	AgentID   string    `json:"agentId"`
	Status    string    `json:"status"`
	Input     string    `json:"input"`
	Output    string    `json:"output,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}