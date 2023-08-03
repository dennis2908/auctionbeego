package controllers

import (
	models "api_beego/models"
	_ "context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/leekchan/accounting"
	_ "github.com/shopspring/decimal"
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

type getBiddings struct {
    Auction_id  int
	Bidder_id int	
	Status int
	Status_name string
}

type cekBiddings struct {
    Auction_id  int
	Bidder_id int	
	Status int
}

type PayBidDownPayment struct {
    Down_payment  int
}


func (api *BiddingsController) GetAllBids() {
    o := orm.NewOrm()
	o.Using("default")
	var Biddings [] ambilBiddings
	var sql string
	sql = "select biddings.*,m_option.name as status_name from biddings "
	sql += "left join m_option on m_option.no = biddings.status and m_option.type = 5"
	fmt.Println(sql)
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
	sql = "select biddings.*,m_option.name as status_name from biddings "
	sql += "left join m_option on m_option.no = biddings.status and m_option.type = 5 where auction_id = '"+api.GetStrings("auction_id")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Biddings)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Biddings	
	}
	
	api.ServeJSON()
}

func AllBidingCheck(api *BiddingsController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	ul := &cekBiddings{}
	json.Unmarshal(frm, ul)

	Auction_id := ul.Auction_id
	Bidder_id := ul.Bidder_id
	Status := ul.Status

	u := &cekBiddings{Auction_id, Bidder_id, Status}
	valid.Required(u.Auction_id, "Auction_id")
	valid.Required(u.Bidder_id, "Bidder_id")
	valid.Required(u.Status, "Status")

	if valid.HasErrors() {
		// If there are error messages it means the validation didn't pass
		// Print error message
		for _, err := range valid.Errors {
			return err.Key + err.Message
		}
	}

	return ""
}

func (api *BiddingsController) CreateBiddings() {
    frm := api.Ctx.Input.RequestBody

	if AllBidingCheck(api) != "" {
		api.Data["json"] = AllBidingCheck(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	u:= &cekBiddings{}
	json.Unmarshal(frm,u)
	Auction_id := u.Auction_id
	Bidder_id := u.Bidder_id
	Status := u.Status
	PostsQry := models.Biddings{Auction_id: Auction_id, Bidder_id: Bidder_id, Status: Status}
	_, err := o.Insert(&PostsQry)
	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully save data"
	api.ServeJSON()
}

func (api *BiddingsController) CancelBid() {
    if api.Ctx.Input.Param(":id") == "" {
        api.Ctx.WriteString("Bid id is empty")
        return
    }

	
    o := orm.NewOrm()
	o.Using("default")
	var Biddings []*models.Biddings
	var sql string
	sql = "select bid_id from biddings where bid_id = '"+api.Ctx.Input.Param(":id")+"'";
	num, err := o.Raw(sql).QueryRows(&Biddings)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE biddings SET status = 3 where bid_id = '"+api.Ctx.Input.Param(":id")+"'" 	
	o.Raw(sql).QueryRows(&Biddings)
	api.Data["json"] = "successfully cancel biddings with bid_id = '"+api.Ctx.Input.Param(":id")+"'" 	 	
	api.ServeJSON()
}

func (api *BiddingsController) PayBidDownPayment() {
	frm := api.Ctx.Input.RequestBody
	u:= &PayBidDownPayment{}
	json.Unmarshal(frm,u)

	Down_payment := strconv.Itoa(u.Down_payment)

    if Down_payment == "" {
        api.Ctx.WriteString("Down payment is empty")
        return
    }
	if api.Ctx.Input.Param(":id") == "" {
        api.Ctx.WriteString("bid id is empty")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Biddings []*models.Biddings
	var sql string
	sql = "select bid_id from biddings where bid_id = '"+api.Ctx.Input.Param(":id")+"'";
	num, err := o.Raw(sql).QueryRows(&Biddings)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = "Failed pay bid"
	    api.ServeJSON()   
	}
	sql = "UPDATE auctions SET down_payment = "+Down_payment+" where auction_id = '"+api.Ctx.Input.Param(":id")+"'" 	
	o.Raw(sql).QueryRows(&Biddings)
	api.Data["json"] = "successfully pay bid with auction_id = "+api.Ctx.Input.Param(":id")
	api.ServeJSON()
}