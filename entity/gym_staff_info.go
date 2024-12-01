package entity

import "time"

type GymStaffInfo struct {
	Id                     uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	GymId                  uint          `json:"gym_id"`
	EnquiryId              uint          `json:"enquiry_id"`
	PlanId                 uint          `json:"plan_id"`
	AddressId              uint          `json:"address_id"`
	AddressInfo            AddressMaster `json:"address_info" gorm:"-"`
	Role                   string        `json:"role"`
	ProfilePicture         string        `json:"profile_picture"`
	FirstName              string        `json:"first_name"`
	LastName               string        `json:"last_name"`
	DateOfBirth            string        `json:"date_of_birth"`
	MobileNo               string        `json:"mobile_no"`
	Email                  string        `json:"email"`
	Gender                 string        `json:"gender"`
	AssignTo               string        `json:"assign_to"`
	EmergencyContactPerson string        `json:"emergency_contact_person"`
	EmergencyContactNo     string        `json:"emergency_contact_no"`
	RegistrationDate       string        `json:"registration_date"`
	MemberGST              string        `json:"member_gst"`
	ReferenceDetails       string        `json:"reference_details"`
	BatchDetails           string        `json:"batch_details"`
	Height                 *float64      `json:"height"`
	Weight                 *float64      `json:"weight"`
	FitnessAspiration      string        `json:"fitness_aspiration"`
	Category               string        `json:"category"`
	Activity               string        `json:"activity"`
	HealthCondition        string        `json:"health_condition"`
	AddOnServices          string        `json:"add_on_services"`
	Specialisation         string        `json:"specialisation"`
	Experience             *int          `json:"experience"`
	TrainerFor             string        `json:"trainer_for"`
	SalaryDetails          *float64      `json:"salary_details"`
	SalaryDeduction        *float64      `json:"salary_deduction"`
	LateMarkAllowPerMonth  *int          `json:"late_mark_allow_per_month"`
	LateComingDeduction    *float64      `json:"late_coming_deduction"`
	CreatedAt              time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt              time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy              *int          `json:"created_by"`
	UpdatedBy              *int          `json:"updated_by"`
}

func (GymStaffInfo) TableName() string {
	return "gym_staff_info"
}
