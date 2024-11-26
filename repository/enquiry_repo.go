package repository

import (
	"errors"
	"fmt"
	"fytrack/config"
	"fytrack/entity"

	"gorm.io/gorm"
)

func AddEnquiryInfo(enquiry entity.Enquiry) error {
	if err := config.DB.Create(&enquiry).Error; err != nil {
		return errors.New("failed to add enquiry to the database" + err.Error())
	}
	return nil
}

func GetEnquiryInfo(enquiryId int64) ([]entity.Enquiry, error) {
	fmt.Printf("Fetching enquiry with ID: %d\n", enquiryId)
	var enquiries []entity.Enquiry

	if enquiryId != 0 {
		if err := config.DB.Where("id = ?", enquiryId).Find(&enquiries).Error; err != nil {
			return nil, errors.New("enquiry not found: " + err.Error())
		}
	} else {
		if err := config.DB.Find(&enquiries).Error; err != nil {
			return nil, errors.New("failed to fetch enquiry: " + err.Error())
		}
	}

	return enquiries, nil
}

func UpdateEnquiryInfo(enquiryId int64, enquiry entity.Enquiry) error {
	var existingEnquiry entity.Enquiry
	if err := config.DB.First(&existingEnquiry, enquiryId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("enquiry not found")
		}
		return errors.New("failed to find existingEnquiry: " + err.Error())
	}
	if err := config.DB.Model(&existingEnquiry).Updates(enquiry).Error; err != nil {
		return errors.New("failed to update existingEnquiry: " + err.Error())
	}
	return nil
}

func DeleteEnquiryInfo(enquiryId int64) error {
	var existingEnquiry entity.Enquiry
	if err := config.DB.First(&existingEnquiry, enquiryId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("enquiry not found")
		}
		return errors.New("failed to find Enquiry: " + err.Error())
	}

	if err := config.DB.Delete(&existingEnquiry).Error; err != nil {
		return errors.New("failed to delete Enquiry: " + err.Error())
	}

	return nil
}
