package models

import (
	"time"
)

type Auctions struct {
	Auction_id int `orm:"pk;auto"`
	Coll_id int
	From_date time.Time `orm:"type(date)"`	
	Due_date time.Time `orm:"type(date)"`
	Auction_method int
	Down_payment int
	Description string
	Auction_count int
	Status int 
	Accepted_bidder int
}

func (a *Auctions) TableName() string {
	return "auctions"
}