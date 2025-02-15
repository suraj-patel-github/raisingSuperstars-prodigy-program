package service

import (
	"prodigy-program/repos"
)


type WeekPlanService interface {
	CreateWeekPlan(description string, userID int) (int, error)
	GetWeekPlanByID(id int) (map[string]interface{}, error)
}

type weekPlanService struct {
	repo repos.WeekPlanRepository
}

func NewWeekPlanService(repo repos.WeekPlanRepository) WeekPlanService {
	return &weekPlanService{repo: repo}
}

func (s *weekPlanService) CreateWeekPlan(description string, userID int) (int, error) {
	return s.repo.CreateWeekPlan(description, userID)
}

func (s *weekPlanService) GetWeekPlanByID(id int) (map[string]interface{}, error) {
	weekPlan, err := s.repo.GetWeekPlanByID(id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":          weekPlan.ID,
		"description": weekPlan.Description,
		"created_at":  weekPlan.CreatedAt,
	}, nil
}