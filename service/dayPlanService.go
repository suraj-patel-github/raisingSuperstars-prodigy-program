package service

import (
	"prodigy-program/repos"
	"prodigy-program/types"
)

type DayPlanService interface {
	UpdateDayPlan(dayNumber int, activities []types.Activity) error
}

type dayPlanService struct{
	repo repos.DayPlanRepository
}

func NewDayPlanService(repo repos.DayPlanRepository) DayPlanService {
	return &dayPlanService{repo: repo}
}

func (s *dayPlanService) UpdateDayPlan(dayNumber int, activities []types.Activity) error {
	return s.repo.UpdateDayPlan(dayNumber, activities)
}