package model

import (
    "time"
)

// Todo represents a single todo item
type Todo struct {
    ID          int64      `json:"id"`                     // Auto-increment or UUID
    Title       string     `json:"title"`                  // Short title or name
    Description string     `json:"description"`            // Optional detailed description
    Completed   bool       `json:"completed"`              // Task status
    DueDate     *time.Time `json:"due_date,omitempty"`     // Optional due date
    Priority    string     `json:"priority"`               // low, medium, high
    CreatedAt   time.Time  `json:"created_at"`             // Timestamp when added
    UpdatedAt   time.Time  `json:"updated_at"`             // Timestamp when last updated
    UserID      int64      `json:"user_id"`                // Foreign key to a User (if multi-user)
}
