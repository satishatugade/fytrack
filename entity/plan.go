package entity

import "time"

type Plan struct {
	Id             int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	PlanName       string    `json:"plan_name"`
	Category       string    `json:"category"`
	PlanType       string    `json:"plan_type"`
	PlanFor        string    `json:"plan_for"`
	ShowPlanOnline string    `json:"show_plan_online"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
}

func (Plan) TableName() string {
	return "plan_details"
}
