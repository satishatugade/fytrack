package repository

import (
	"errors"
	"fytrack/config"
	"fytrack/entity"
)

func AddMemberAddressInfo(address entity.AddressMaster) (int64, error) {
	if err := config.DB.Create(&address).Error; err != nil {
		return 0, err
	}
	return address.Id, nil
}

func GetAddressMasterInfo(addressID int) (entity.AddressMaster, error) {
	var address entity.AddressMaster
	if err := config.DB.Where("id = ?", addressID).First(&address).Error; err != nil {
		return address, errors.New("failed to fetch address info: " + err.Error())
	}
	return address, nil
}
