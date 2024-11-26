package entity

import "time"

type Enquiry struct {
	ID                  uint          `gorm:"primaryKey" json:"id"`
	AddressId           uint          `json:"address_id"`
	AddressInfo         AddressMaster `json:"address_info" gorm:"-"`
	FirstName           *string       `json:"first_name"`
	LastName            *string       `json:"last_name"`
	MobileNo            *string       `json:"mobile_no"`
	Gender              *string       `json:"gender"`
	IntrestedSports     *string       `json:"intrested_sports"`
	EnquiryDate         *string       `json:"enquiry_date"`
	EnquiryStatus       *string       `json:"enquiry_status"`
	ExpectedJoiningDate *string       `json:"expected_joining_date"`
	LeadType            *string       `json:"lead_type"`
	Remark              *string       `json:"remark"`
	LastUpdatedDate     *string       `json:"last_updated_date"`
	LastCallStatus      *string       `json:"last_call_status"`
	FollowUpDate        *string       `json:"follow_up_date"`
	ReferenceType       *string       `json:"reference_type"`
	Reference           *string       `json:"reference"`
	AssignTo            *string       `json:"assign_to"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	CreatedBy           *string       `json:"created_by"`
	UpdatedBy           *string       `json:"updated_by"`
}

func (Enquiry) TableName() string {
	return "enquiry_data"
}
