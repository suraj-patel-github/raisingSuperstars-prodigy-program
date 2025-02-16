package repos

import (
	"database/sql"
	"fmt"
	"prodigy-program/types"
	"time"
)

type WeekPlanRepository interface {
	CreateWeekPlan(description string, userID int) (int, error)
	GetWeekPlan(weekID int, dayNumber *int) (*types.WeekPlanResponse, error)
}

type weekPlanRepo struct {
	db *sql.DB
}

func NewWeekPlanRepository(db *sql.DB) WeekPlanRepository {
	return &weekPlanRepo{db: db}
}

// Insert new week plan and populate related tables
func (r *weekPlanRepo) CreateWeekPlan(description string, userID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// Insert into WeekPlan
	var weekID int
	err = tx.QueryRow(`INSERT INTO weekplan (description) VALUES ($1) RETURNING id`, description).Scan(&weekID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Get all activity IDs
	rows, err := tx.Query(`SELECT id FROM activitydesc`)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer rows.Close()

	var activityIDs []int
	for rows.Next() {
		var activityID int
		if err := rows.Scan(&activityID); err != nil {
			tx.Rollback()
			return 0, err
		}
		activityIDs = append(activityIDs, activityID)
	}

	// Ensure there are at least 9 activities
	if len(activityIDs) < 9 {
		tx.Rollback()
		return 0, fmt.Errorf("not enough activities found")
	}

	// Assign 9 activities per day for 7 days
	for day := 0; day < 7; day++ {
		for i := 0; i < 9; i++ {
			_, err := tx.Exec(
				`INSERT INTO dayplan (user_id, week_id, day_number, activity_id, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`,
				userID, weekID, day, activityIDs[i],
			)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return weekID, nil
}

// Get week plan by ID
func (r *weekPlanRepo) GetWeekPlan(weekID int, dayNumber *int) (*types.WeekPlanResponse, error) {
	// Base Query
	query := `
		SELECT dp.id, dp.user_id, dp.week_id, dp.day_number, dp.activity_id, dp.completed_at, dp.created_at, 
		       a.category, a.activity_name, a.time, a.frequency
		FROM dayplan dp
		JOIN activitydesc a ON dp.activity_id = a.id
		WHERE dp.week_id = $1
	`

	// If dayNumber is provided, filter by it
	var rows *sql.Rows
	var err error
	if dayNumber != nil {
		query += " AND dp.day_number = $2 ORDER BY dp.created_at"
		rows, err = r.db.Query(query, weekID, *dayNumber)
	} else {
		query += " ORDER BY dp.created_at"
		rows, err = r.db.Query(query, weekID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Prepare response
	weekPlan := &types.WeekPlanResponse{
		WeekID:      weekID,
		Description: "description",
		DayPlans:    make(map[int][]types.Activity),
	}

	// Iterate over results
	for rows.Next() {
		var activity types.Activity
		var createdAt time.Time
		var completedAt sql.NullTime

		err := rows.Scan(
			&activity.ID, &activity.UserID, &activity.WeekID, &activity.DayNumber, &activity.ActivityID,
			&completedAt, &createdAt,
			&activity.Category, &activity.ActivityName, &activity.Time, &activity.Frequency,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable completed_at
		if completedAt.Valid {
			activity.CompletedAt = completedAt.Time
		} else {
			activity.CompletedAt = time.Time{}
		}

		// Append to the correct day's activities
		weekPlan.DayPlans[activity.DayNumber] = append(weekPlan.DayPlans[activity.DayNumber], activity)
	}

	return weekPlan, nil
}
