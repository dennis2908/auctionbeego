package models

type Biddings struct {
	Bid_id int `orm:"pk;auto"`
	Auction_id  int
	Bidder_id int	
	Status int
}

func (a *Biddings) TableName() string {
	return "biddings"
}