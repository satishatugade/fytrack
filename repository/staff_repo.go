package repository

import (
	"errors"
	"fmt"
	"fytrack/config"
	"fytrack/entity"

	"gorm.io/gorm"
)

func AddStaffInfo(staff entity.GymStaffInfo) error {
	if err := config.DB.Create(&staff).Error; err != nil {
		return errors.New("failed to add staff to the database" + err.Error())
	}
	return nil
}

func GetStaffInfo(staffId int64) ([]entity.GymStaffInfo, error) {
	fmt.Printf("Fetching staff with ID: %d\n", staffId)
	var staffs []entity.GymStaffInfo

	if staffId != 0 {
		if err := config.DB.Where("id = ?", staffId).Find(&staffs).Error; err != nil {
			return nil, errors.New("staff not found: " + err.Error())
		}
	} else {
		if err := config.DB.Find(&staffs).Error; err != nil {
			return nil, errors.New("failed to fetch staffs: " + err.Error())
		}
	}

	return staffs, nil
}

func UpdateStaffInfo(staffId int64, staff entity.GymStaffInfo) (entity.GymStaffInfo, error) {
	var existingstaff entity.GymStaffInfo

	if err := config.DB.First(&existingstaff, staffId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.GymStaffInfo{}, errors.New("staff not found")
		}
		return entity.GymStaffInfo{}, errors.New("failed to find member: " + err.Error())
	}

	if err := config.DB.Model(&existingstaff).Updates(staff).Error; err != nil {
		return entity.GymStaffInfo{}, errors.New("failed to update staff: " + err.Error())
	}
	if err := config.DB.First(&existingstaff, staffId).Error; err != nil {
		return entity.GymStaffInfo{}, errors.New("failed to reload updated staff: " + err.Error())
	}

	return existingstaff, nil
}

func DeleteStaffInfo(staffId int64) error {
	var existingstaff entity.GymStaffInfo
	if err := config.DB.First(&existingstaff, staffId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("staff not found")
		}
		return errors.New("failed to find staff: " + err.Error())
	}
	if err := config.DB.Delete(&existingstaff).Error; err != nil {
		return errors.New("failed to delete staff: " + err.Error())
	}

	return nil
}
