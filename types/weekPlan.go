package types

import "time"

type WeekPlan struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type DayPlan struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	WeekID      int        `json:"week_id"`
	ActivityID  int        `json:"activity_id"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
