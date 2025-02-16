package service

import (
	"prodigy-program/repos"
	"prodigy-program/types"
)


type WeekPlanService interface {
	CreateWeekPlan(description string, userID int) (int, error)
	GetWeekPlanByID(weekID int, dayNumber *int) (*types.WeekPlanResponse, error)
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

func (s *weekPlanService) GetWeekPlanByID(weekID int, dayNumber *int) (*types.WeekPlanResponse, error) {
	return s.repo.GetWeekPlan(weekID, dayNumber)
}