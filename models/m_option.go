package models

type M_option struct {
	Id int `orm:"pk;auto"`
	Type int	
	No int
	Name string
}

func (a *M_option) TableName() string {
	return "m_option"
}