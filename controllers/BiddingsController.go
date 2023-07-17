package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/shopspring/decimal"
    _ "github.com/leekchan/accounting"
	_ "github.com/astaxie/beego/validation"
	models "api_beego/models"
	_ "strconv"
)

type BiddingsController struct {
	beego.Controller
}

type ambilBiddings struct {
    Bid_id int
	Auction_id  int
	Bidder_id int	
	Status int
	Status_name string
}


func (api *BiddingsController) GetAllBids() {
    o := orm.NewOrm()
	o.Using("default")
	var Biddings [] ambilBiddings
	var sql string
	sql = "select biddings.*,m_option.name as status_name from biddings"
	sql += " join m_option on m_option.no = biddings.status and m_option.type = 5"
	num, err := o.Raw(sql).QueryRows(&Biddings)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Biddings		
	}
	
	api.ServeJSON()
}

func (api *BiddingsController) GetBidsByAuction() {
    auction_id := api.GetString("auction_id")
    if auction_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Biddings [] ambilBiddings
	var sql string
	sql = "select biddings.*,m_option.name as status_name from biddings"
	sql += " join m_option on m_option.no = biddings.status and m_option.type = 5 where auction_id = '"+api.GetStrings("auction_id")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Biddings)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Biddings[0]		
	}
	
	api.ServeJSON()
}

func (api *BiddingsController) CreateBiddings() {
    auction_id := api.GetString("auction_id")
    if auction_id == "" {
        api.Ctx.WriteString("")
        return
    }
	bidder_id := api.GetString("bidder_id")
    if bidder_id == "" {
        api.Ctx.WriteString("")
        return
    }
	status := api.GetString("status")
    if status == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Biddings []*models.Biddings
	var sql string
	sql = "INSERT INTO biddings (auction_id,bidder_id,status) VALUES ('"+api.GetStrings("auction_id")[0]+"','"+api.GetStrings("bidder_id")[0]+"','"+api.GetStrings("status")[0]+"')"
	o.Raw(sql).QueryRows(&Biddings)
	api.Data["json"] = 1
	api.ServeJSON()
}

func (api *BiddingsController) CancelBid() {
    bid_id := api.GetString("bid_id")
    if bid_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Biddings []*models.Biddings
	var sql string
	sql = "select bid_id from biddings where bid_id = '"+api.GetString("bid_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Biddings)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE biddings SET status = 3 where bid_id = '"+api.GetStrings("bid_id")[0]+"'" 	
	o.Raw(sql).QueryRows(&Biddings)
	api.Data["json"] = "successfully cancel biddings with bid_id = "+api.GetStrings("bid_id")[0]
	api.ServeJSON()
}

func (api *BiddingsController) PayBidDownPayment() {
    down_payment := api.GetString("down_payment")
    if down_payment == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Biddings []*models.Biddings
	var sql string
	sql = "select bid_id from biddings where bid_id = '"+api.GetString("bid_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Biddings)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE auctions SET down_payment = '"+api.GetStrings("down_payment")[0]+"' where auction_id = '"+api.GetStrings("auction_id")[0]+"'" 	
	o.Raw(sql).QueryRows(&Biddings)
	api.Data["json"] = "successfully pay bid with auction_id = "+api.GetStrings("auction_id")[0]
	api.ServeJSON()
}