package controller

import (
	"fytrack/entity"
	"fytrack/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddStaffInfo(c *gin.Context) {
	var staff entity.GymStaffInfo
	if err := c.ShouldBindJSON(&staff); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}
	AddressId, AddressErr := service.AddressMaster(staff.AddressInfo)
	if AddressErr != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to add staff address info", AddressErr.Error())
		return
	}
	staff.AddressId = uint(AddressId)
	err := service.AddStaffInfo(staff)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to add staff info", err.Error())
		return
	}
	entity.RespondSuccess(c, "staff added successfully", nil)
}

func GetStaffInfo(c *gin.Context) {
	staffIdStr := c.Query("id")
	var staffId int64
	var err error
	if staffIdStr != "" {
		staffId, err = strconv.ParseInt(staffIdStr, 10, 64)
		if err != nil {
			entity.RespondError(c, http.StatusBadRequest, "Invalid staff Id", err.Error())
			return
		}
	}

	staffs, err := service.GetStaffInfo(staffId)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to fetch staff info ", err.Error())
		return
	}
	for i := range staffs {
		AddressInfo, AddressErr := service.GetAddressMasterInfo(staffs[i].AddressId)
		if err != nil {
			entity.RespondError(c, http.StatusInternalServerError, "Unable to fetch staff address info ", AddressErr.Error())
			return
		}
		staffs[i].AddressInfo = AddressInfo
	}
	entity.RespondSuccess(c, "staff information retrieved successfully", staffs)
}

func UpdateStaffInfo(c *gin.Context) {
	staffIdStr := c.Query("id")
	staffId, err := strconv.ParseInt(staffIdStr, 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid member ID : ", err.Error())
		return
	}
	var staff entity.GymStaffInfo
	if err := c.ShouldBindJSON(&staff); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data ", err.Error())
		return
	}
	updatedStaffInfo, err := service.UpdateStaffInfo(staffId, staff)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to update staff info ", err.Error())
		return
	}
	err = service.UpdateAddressMasterInfo(updatedStaffInfo.AddressId, staff.AddressInfo)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to update staff address info ", err.Error())
		return
	}

	entity.RespondSuccess(c, "staff info updated successfully", nil)
}

func DeleteStaffInfo(c *gin.Context) {
	staffIdStr := c.Query("id")
	staffId, err := strconv.ParseInt(staffIdStr, 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid staff Id ", err.Error())
		return
	}

	err = service.DeleteStaffInfo(staffId)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to delete staff info: ", err.Error())
		return
	}
	err = service.DeleteAddressInfo(staffId)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to delete staff address info: ", err.Error())
		return
	}

	entity.RespondSuccess(c, "staff info deleted successfully", nil)
}
