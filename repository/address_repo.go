package repository

import (
	"errors"
	"fytrack/config"
	"fytrack/entity"

	"gorm.io/gorm"
)

func AddMemberAddressInfo(address entity.AddressMaster) (int64, error) {
	if err := config.DB.Create(&address).Error; err != nil {
		return 0, err
	}
	return address.Id, nil
}

func GetAddressMasterInfo(addressID uint) (entity.AddressMaster, error) {
	var address entity.AddressMaster
	if err := config.DB.Where("id = ?", addressID).First(&address).Error; err != nil {
		return address, errors.New("failed to fetch address info: " + err.Error())
	}
	return address, nil
}

func UpdateAddressMasterInfo(addressId uint, addressInfo entity.AddressMaster) error {
	var existingAddress entity.AddressMaster
	if err := config.DB.First(&existingAddress, addressId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("existing Address not found")
		}
		return errors.New("failed to find address : " + err.Error())
	}

	if err := config.DB.Model(&existingAddress).Updates(addressInfo).Error; err != nil {
		return errors.New("failed to update address info : " + err.Error())
	}

	return nil
}

func DeleteAddressInfo(addressId int64) error {
	var existingAddressInfo entity.AddressMaster
	if err := config.DB.First(&existingAddressInfo, addressId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("address not found")
		}
		return errors.New("failed to find address info: " + err.Error())
	}

	if err := config.DB.Delete(&existingAddressInfo).Error; err != nil {
		return errors.New("failed to delete address info : " + err.Error())
	}

	return nil
}
