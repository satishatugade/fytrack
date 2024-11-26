package controller

import (
	"fytrack/entity"
	"fytrack/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMemberData(c *gin.Context) {
	var member entity.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}
	AddressId, AddressErr := service.AddressMaster(member.AddressInfo)
	if AddressErr != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to add member address info", AddressErr.Error())
		return
	}
	member.AddressId = uint(AddressId)
	err := service.AddMemberInfo(member)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to add member info", err.Error())
		return
	}
	entity.RespondSuccess(c, "Member added successfully", nil)
}

func GetMemberInfo(c *gin.Context) {
	memberIDStr := c.Query("id")
	var memberID int64
	var err error
	if memberIDStr != "" {
		memberID, err = strconv.ParseInt(memberIDStr, 10, 64)
		if err != nil {
			entity.RespondError(c, http.StatusBadRequest, "Invalid member Id", err.Error())
			return
		}
	}

	members, err := service.GetMemberInfo(memberID)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to fetch member info ", err.Error())
		return
	}
	for i := range members {
		AddressInfo, AddressErr := service.GetAddressMasterInfo(members[i].AddressId)
		if err != nil {
			entity.RespondError(c, http.StatusInternalServerError, "Unable to fetch member address info ", AddressErr.Error())
			return
		}
		members[i].AddressInfo = AddressInfo
	}
	entity.RespondSuccess(c, "Member information retrieved successfully", members)
}

func UpdateMemberInfo(c *gin.Context) {
	memberIDStr := c.Query("id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid member ID : ", err.Error())
		return
	}
	var member entity.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data ", err.Error())
		return
	}
	err = service.UpdateMemberInfo(memberID, member)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to update member info ", err.Error())
		return
	}
	err = service.UpdateAddressMasterInfo(uint(memberID), member.AddressInfo)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to update address info ", err.Error())
		return
	}

	entity.RespondSuccess(c, "Member info updated successfully", nil)
}

func DeleteMemberInfo(c *gin.Context) {
	memberIDStr := c.Query("id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid member Id ", err.Error())
		return
	}

	err = service.DeleteMemberInfo(memberID)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to delete member info: ", err.Error())
		return
	}
	err = service.DeleteAddressInfo(memberID)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to delete address info: ", err.Error())
		return
	}

	entity.RespondSuccess(c, "Member info deleted successfully", nil)
}
