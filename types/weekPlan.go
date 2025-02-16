package types

import "time"

type WeekPlan struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type DayPlan struct {
	ID          int        `json:"id"`
	UserID      int        `json:"userId"`
	WeekID      int        `json:"weekId"`
	ActivityID  int        `json:"activityId"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type WeekPlanResponse struct {
	WeekID      int                `json:"weekId"`
	Description string             `json:"description"`
	DayPlans    map[int][]Activity `json:"days"`
}

type Activity struct {
	ID           int       `json:"id,omitempty"`
	UserID       int       `json:"userId"`
	WeekID       int       `json:"weekId"`
	ActivityID   int       `json:"activityId"`
	DayNumber    int       `json:"dayNumber"`
	Category     string    `json:"category"`
	ActivityName string    `json:"activity_name"`
	Time         string    `json:"time"`
	Frequency    string    `json:"frequency"`
	CompletedAt  time.Time `json:"completedAt,omitempty"`
}
