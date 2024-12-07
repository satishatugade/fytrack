package repository

import (
	"errors"
	"fmt"
	"fytrack/config"
	"fytrack/entity"

	"gorm.io/gorm"
)

func AddPlanInfo(plan entity.Plan) error {
	if err := config.DB.Create(&plan).Error; err != nil {
		return err
	}
	return nil
}

func GetPlanInfo(planId int64) ([]entity.Plan, error) {
	fmt.Printf("Fetching plan with ID: %d\n", planId)
	var plans []entity.Plan

	if planId != 0 {
		if err := config.DB.Where("id = ?", planId).Find(&plans).Error; err != nil {
			return nil, errors.New("plan not found: " + err.Error())
		}
	} else {
		if err := config.DB.Find(&plans).Error; err != nil {
			return nil, errors.New("failed to fetch plan: " + err.Error())
		}
	}

	return plans, nil
}

func UpdatePlanInfo(planId int64, updatedPlan entity.Plan) (entity.Plan, error) {
	var plan entity.Plan
	if err := config.DB.First(&plan, planId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Plan{}, errors.New("plan not found")
		}
		return entity.Plan{}, err
	}

	if err := config.DB.Model(&plan).Updates(updatedPlan).Error; err != nil {
		return entity.Plan{}, err
	}

	return plan, nil
}

func DeletePlanInfo(planId int64) error {
	if err := config.DB.Delete(&entity.Plan{}, planId).Error; err != nil {
		return err
	}
	return nil
}
