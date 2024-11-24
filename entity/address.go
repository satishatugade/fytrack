package entity

type AddressMaster struct {
	Id        int64  `json:"id"`
	GymId     int64  `json:"gym_id"`
	EnquiryId int64  `json:"enquiry_id"`
	Address1  int64  `json:"address1"`
	Address2  string `json:"address2"`
	Locality  string `json:"locality"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Pincode   string `json:"pincode"`
}

func (AddressMaster) TableName() string {
	return "address_master"
}
