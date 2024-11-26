package controller

import (
	"fytrack/entity"
	"fytrack/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddEnquiryData(c *gin.Context) {
	var enquiry entity.Enquiry
	if err := c.ShouldBindJSON(&enquiry); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}
	AddressId, AddressErr := service.AddressMaster(enquiry.AddressInfo)
	if AddressErr != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to add member address info", AddressErr.Error())
		return
	}
	enquiry.AddressId = uint(AddressId)
	err := service.AddEnquiryInfo(enquiry)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to add member info", err.Error())
		return
	}
	entity.RespondSuccess(c, "Client enquiry added successfully", nil)
}

func GetEnquiryInfo(c *gin.Context) {
	enquiryIdStr := c.Query("id")
	var enquiryId int64
	var err error
	if enquiryIdStr != "" {
		enquiryId, err = strconv.ParseInt(enquiryIdStr, 10, 64)
		if err != nil {
			entity.RespondError(c, http.StatusBadRequest, "Invalid enquiry Id", err.Error())
			return
		}
	}

	enquiries, err := service.GetEnquiryInfo(enquiryId)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to fetch enquiry info ", err.Error())
		return
	}
	for i := range enquiries {
		AddressInfo, AddressErr := service.GetAddressMasterInfo(enquiries[i].AddressId)
		if err != nil {
			entity.RespondError(c, http.StatusInternalServerError, "Unable to fetch member enquiry address info ", AddressErr.Error())
			return
		}
		enquiries[i].AddressInfo = AddressInfo
	}
	entity.RespondSuccess(c, "Enquiry information retrieved successfully", enquiries)
}

func UpdateEnquiryInfo(c *gin.Context) {
	enquiryIdStr := c.Query("id")
	enquiryId, err := strconv.ParseInt(enquiryIdStr, 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid enquiry Id : ", err.Error())
		return
	}
	var enquiry entity.Enquiry
	if err := c.ShouldBindJSON(&enquiry); err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid request data ", err.Error())
		return
	}
	err = service.UpdateEnquiryInfo(enquiryId, enquiry)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to update enquiry info ", err.Error())
		return
	}
	err = service.UpdateAddressMasterInfo(uint(enquiryId), enquiry.AddressInfo)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to update enquiry address info ", err.Error())
		return
	}

	entity.RespondSuccess(c, "Enquiry info updated successfully", nil)
}

func DeleteEnquiryInfo(c *gin.Context) {
	enquiryIdStr := c.Query("id")
	enquiryId, err := strconv.ParseInt(enquiryIdStr, 10, 64)
	if err != nil {
		entity.RespondError(c, http.StatusBadRequest, "Invalid enquiry Id ", err.Error())
		return
	}

	err = service.DeleteEnquiryInfo(enquiryId)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to delete enquiry info: ", err.Error())
		return
	}
	err = service.DeleteAddressInfo(enquiryId)
	if err != nil {
		entity.RespondError(c, http.StatusInternalServerError, "Unable to delete enquiry address info: ", err.Error())
		return
	}

	entity.RespondSuccess(c, "Enquiry info deleted successfully", nil)
}
