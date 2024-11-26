package entity

type Member struct {
	Id                     uint          `json:"id" gorm:"primaryKey"`
	PlanId                 uint          `json:"plan_id"`
	AddressId              uint          `json:"address_id"`
	AddressInfo            AddressMaster `json:"address_info" gorm:"-"`
	FirstName              string        `json:"first_name" binding:"required"`
	LastName               string        `json:"last_name"`
	Email                  string        `json:"email"`
	MobileNo               string        `json:"mobile_no"`
	AssignTo               string        `json:"assign_to"`
	DateOfBirth            string        `json:"date_of_birth"`
	Gender                 string        `json:"gender"`
	RegistrationDate       string        `json:"registration_date"`
	EmergencyContactPerson string        `json:"emergency_contact_person"`
	EmergencyContact       string        `json:"emergency_contact"`
	ProfilePicture         string        `json:"profile_picture"`
	ReferenceDetails       string        `json:"reference_details"`
	BatchDetails           string        `json:"batch_details"`
	Height                 string        `json:"height"`
	Weight                 string        `json:"weight"`
	FitnessAspiration      string        `json:"fitness_aspiration"`
	Category               string        `json:"category"`
	Activities             string        `json:"activities"`
	HealthCondition        string        `json:"health_condition"`
	BloodGroup             string        `json:"blood_group"`
}

func (Member) TableName() string {
	return "gym_member_info"
}
