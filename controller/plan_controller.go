package controller

import (
	"fytrack/entity"
	"fytrack/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPlanInfo(c *gin.Context) {
	var plan entity.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	err := service.AddPlanInfo(plan)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to add plan info", err.Error())
		return
	}
	entity.RespondSuccess(c, "Plan added successfully", nil)
}

func GetPlanInfo(c *gin.Context) {
	planIdStr := c.Query("id")
	var planId int64
	var err error
	if planIdStr != "" {
		planId, err = strconv.ParseInt(planIdStr, 10, 64)
		if err != nil {
			entity.RespondError(c, http.StatusBadRequest, "Invalid plan Id", err.Error())
			return
		}
	}
	plans, err := service.GetPlanInfo(planId)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to fetch plans info ", err.Error())
		return
	}
	entity.RespondSuccess(c, "Plans information retrieved successfully", plans)
}

func UpdatePlanInfo(c *gin.Context) {
	planIdStr := c.Query("id")
	planId, err := strconv.ParseInt(planIdStr, 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid plan Id : ", err.Error())
		return
	}
	var plan entity.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data ", err.Error())
		return
	}
	updatedPlan, err := service.UpdatePlanInfo(planId, plan)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to update plan info ", err.Error())
		return
	}

	entity.RespondSuccess(c, "Plan updated successfully ", updatedPlan)
}

func DeletePlanInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid plan Id ", err.Error())
		return
	}
	err = service.DeletePlanInfo(id)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to delete plan info : ", err.Error())
		return
	}
	entity.RespondSuccess(c, "Plan deleted successfully", nil)
}
