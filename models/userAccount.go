package models

type UserAccount struct {
	User_id int `orm:"pk;auto"`
	Username string `orm:"unique"`
	Password string	
	Email string
	Status int
	Ktp_no string
}

func (a *UserAccount) TableName() string {
	return "useraccount"
}