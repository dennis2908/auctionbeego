package models

type PartnerAccount struct {
	Partner_id int `orm:"pk;auto"`
	Username string `orm:"unique"`
	Password string
	Email string
	Ljk_id int	
	Status int
	Ktp_no string
}

func (a *PartnerAccount) TableName() string {
	return "partneraccount"
}