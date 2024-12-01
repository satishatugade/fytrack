package service

import (
	"fytrack/entity"
	"fytrack/repository"
)

func AddMemberInfo(member entity.Member) error {
	return repository.AddMemberInfo(member)
}

func GetMemberInfo(memberID int64) ([]entity.Member, error) {
	return repository.GetMemberInfo(memberID)
}

func UpdateMemberInfo(memberId int64, member entity.Member) (entity.Member, error) {
	return repository.UpdateMemberInfo(memberId, member)
}

func DeleteMemberInfo(memberID int64) error {
	return repository.DeleteMemberInfo(memberID)
}
