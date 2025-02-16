package repos

import (
	"database/sql"
	"prodigy-program/types"
)


type DayPlanRepository interface {
	UpdateDayPlan(dayNumber int, activities []types.Activity) error
}

type dayPlanRepo struct{
	db *sql.DB
}

func NewDayPlanRepoitory(db *sql.DB) DayPlanRepository{
	return &dayPlanRepo{db: db}
}

func (r *dayPlanRepo) UpdateDayPlan(dayNumber int, activities []types.Activity) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Prepare statement for updates
	stmt, err := tx.Prepare(`
		UPDATE dayplan 
		SET completed_at = $1, updated_at = NOW()
		WHERE id = $2 AND completed_at IS DISTINCT FROM $1
	`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, activity := range activities {
		_, err := stmt.Exec(activity.CompletedAt, activity.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

