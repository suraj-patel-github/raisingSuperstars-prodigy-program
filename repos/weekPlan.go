package repos

import (
	"database/sql"
	"fmt"
	"prodigy-program/types"
)

type WeekPlanRepository interface {
	CreateWeekPlan(description string, userID int) (int, error)
	GetWeekPlanByID(id int) (*types.WeekPlan, error)
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
				`INSERT INTO dayplan (userId, weekId, activityId, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`,
				userID, weekID, activityIDs[i],
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
func (r *weekPlanRepo) GetWeekPlanByID(id int) (*types.WeekPlan, error) {
	var plan types.WeekPlan
	err := r.db.QueryRow(`SELECT id, description, created_at FROM weekplan WHERE id = $1`, id).
		Scan(&plan.ID, &plan.Description, &plan.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &plan, nil
}