package service

import (
	"fytrack/entity"
	"fytrack/repository"
)

func AddPlanInfo(plan entity.Plan) error {
	return repository.AddPlanInfo(plan)
}

func GetPlanInfo(planId int64) ([]entity.Plan, error) {
	return repository.GetPlanInfo(planId)
}

func UpdatePlanInfo(planId int64, plan entity.Plan) (entity.Plan, error) {
	return repository.UpdatePlanInfo(planId, plan)
}

func DeletePlanInfo(planId int64) error {
	return repository.DeletePlanInfo(planId)
}
