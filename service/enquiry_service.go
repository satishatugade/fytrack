package service

import (
	"fytrack/entity"
	"fytrack/repository"
)

func AddEnquiryInfo(enquiry entity.Enquiry) error {
	return repository.AddEnquiryInfo(enquiry)
}

func GetEnquiryInfo(enquiryId int64) ([]entity.Enquiry, error) {
	return repository.GetEnquiryInfo(enquiryId)
}

func UpdateEnquiryInfo(enquiryId int64, enquiry entity.Enquiry) (entity.Enquiry, error) {
	return repository.UpdateEnquiryInfo(enquiryId, enquiry)
}

func DeleteEnquiryInfo(enquiryId int64) error {
	return repository.DeleteEnquiryInfo(enquiryId)
}
