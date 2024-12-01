package service

import (
	"fytrack/entity"
	"fytrack/repository"
)

func AddStaffInfo(staff entity.GymStaffInfo) error {
	return repository.AddStaffInfo(staff)
}

func GetStaffInfo(staffId int64) ([]entity.GymStaffInfo, error) {
	return repository.GetStaffInfo(staffId)
}

func UpdateStaffInfo(staffId int64, staff entity.GymStaffInfo) (entity.GymStaffInfo, error) {
	return repository.UpdateStaffInfo(staffId, staff)
}

func DeleteStaffInfo(staffId int64) error {
	return repository.DeleteStaffInfo(staffId)
}
