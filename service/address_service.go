package service

import (
	"fytrack/entity"
	"fytrack/repository"
)

func AddressMaster(addressInfo entity.AddressMaster) (int64, error) {
	return repository.AddMemberAddressInfo(addressInfo)
}

func GetAddressMasterInfo(addressId uint) (entity.AddressMaster, error) {
	return repository.GetAddressMasterInfo(addressId)
}

func UpdateAddressMasterInfo(addressId uint, addressInfo entity.AddressMaster) error {
	return repository.UpdateAddressMasterInfo(addressId, addressInfo)
}

func DeleteAddressInfo(addressId int64) error {
	return repository.DeleteAddressInfo(addressId)
}
