package models

import (
	"time"
)

type Logging struct {
	Trx_id int `orm:"pk"`
	Trx_category int
	Details int	
	Timestamp time.Time
}

func (a *Logging) TableName() string {
	return "logging"
}