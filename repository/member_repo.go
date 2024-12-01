package repository

import (
	"errors"
	"fmt"
	"fytrack/config"
	"fytrack/entity"

	"gorm.io/gorm"
)

func AddMemberInfo(member entity.Member) error {
	if err := config.DB.Create(&member).Error; err != nil {
		return errors.New("failed to add member to the database" + err.Error())
	}
	return nil
}

func GetMemberInfo(memberID int64) ([]entity.Member, error) {
	fmt.Printf("Fetching member with ID: %d\n", memberID)
	var members []entity.Member

	if memberID != 0 {
		if err := config.DB.Where("id = ?", memberID).Find(&members).Error; err != nil {
			return nil, errors.New("member not found: " + err.Error())
		}
	} else {
		if err := config.DB.Find(&members).Error; err != nil {
			return nil, errors.New("failed to fetch members: " + err.Error())
		}
	}

	return members, nil
}

func UpdateMemberInfo(memberId int64, member entity.Member) (entity.Member, error) {
	var existingMember entity.Member

	if err := config.DB.First(&existingMember, memberId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Member{}, errors.New("member not found")
		}
		return entity.Member{}, errors.New("failed to find member: " + err.Error())
	}

	if err := config.DB.Model(&existingMember).Updates(member).Error; err != nil {
		return entity.Member{}, errors.New("failed to update member: " + err.Error())
	}
	if err := config.DB.First(&existingMember, memberId).Error; err != nil {
		return entity.Member{}, errors.New("failed to reload updated staff: " + err.Error())
	}
	return existingMember, nil
}

func DeleteMemberInfo(memberID int64) error {
	var existingMember entity.Member
	if err := config.DB.First(&existingMember, memberID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("member not found")
		}
		return errors.New("failed to find member: " + err.Error())
	}

	if err := config.DB.Delete(&existingMember).Error; err != nil {
		return errors.New("failed to delete member: " + err.Error())
	}

	return nil
}
