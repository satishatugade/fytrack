package service

import (
	"fytrack/entity"
	"fytrack/repository"
)

func AddressMaster(address entity.AddressMaster) (int64, error) {
	return repository.AddMemberAddressInfo(address)
}

func GetAddressMasterInfo(addressID int) (entity.AddressMaster, error) {
	return repository.GetAddressMasterInfo(addressID)
}
