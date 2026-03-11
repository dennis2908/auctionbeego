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
	"time"
)

type AuctionsController struct {
	beego.Controller
}

type ambilAuctions struct {
    Auction_id int
	Coll_id int
	From_date time.Time
	Due_date time.Time
	Auction_method int
	Down_payment int
	Description string
	Auction_count int
	Status int 
	Accepted_bidder int
	Acc_id  int
	Type_id int	
	Owner_name string
	Coll_location string
	Initial_col_price int 
	Final_col_price int
	Ljk_id int
	Document_path string
	Username string
	Email string
	Type_name string
	Auction_method_name string
	Accepted_bidder_name string
}


type getAuctions struct {
	Coll_id int
	From_date string
	Due_date string
	Auction_method int
	Down_payment int
	Description string
	Auction_count int
	Status int 
}

type getAuctionsWithId struct {
	Auction_id int
	Coll_id int
	From_date string
	Due_date string
	Auction_method int
	Down_payment int
	Description string
	Auction_count int
	Status int 
}

func (api *AuctionsController) GetAllAuctions() {
    o := orm.NewOrm()
	o.Using("default")
	var Auctions [] ambilAuctions
	var sql string
	sql = "select auctions.*,collateral.*,option_method.name as auction_method_name,useraccount.username as accepted_bidder_name from auctions"
	sql += " left join m_option as option_method on auctions.auction_method  = option_method.no and option_method.type = 2"
	sql += " left join collateral on auctions.coll_id  = collateral.coll_id"
	sql += " left join useraccount on useraccount.user_id  = auctions.accepted_bidder"
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Auctions	
	}
	
	api.ServeJSON()
}

func (api *AuctionsController) GetAuctionsByID() {
    chk := api.GetString("auction_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Auctions [] ambilAuctions
	var sql string
	sql = "select auctions.*,collateral.*,option_method.name as auction_method_name,useraccount.username as accepted_bidder_name from auctions"
	sql += " left join m_option as option_method on option_method.no  = auctions.auction_method and option_method.type = 2"
	sql += " left join collateral on auctions.coll_id  = collateral.coll_id"
	sql += " left join useraccount on useraccount.user_id  = auctions.accepted_bidder where auctions.auction_id = '"+api.GetStrings("auction_id")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Auctions[0]		
	}
	api.ServeJSON()
}

func (api *AuctionsController) GetAuctionsByCollateral() {
    chk := api.GetString("coll_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Auctions [] ambilAuctions
	var sql string
	sql = "select auctions.*,collateral.*,option_method.name as auction_method_name,useraccount.username as accepted_bidder_name from auctions"
	sql += " left join m_option as option_method on option_method.no  = auctions.auction_method and option_method.type = 2"
	sql += " left join collateral on auctions.coll_id  = collateral.coll_id"
	sql += " left join useraccount on useraccount.user_id  = auctions.accepted_bidder where auctions.coll_id = '"+api.GetStrings("coll_id")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Auctions	
	}
	api.ServeJSON()
}

func (api *AuctionsController) GetAuctionsByDueDate() {
    chk := api.GetString("due_date")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Auctions [] ambilAuctions
	var sql string
	sql = "select auctions.*,collateral.*,option_method.name as auction_method_name,useraccount.username as accepted_bidder_name from auctions"
	sql += " left join m_option as option_method on option_method.no  = auctions.auction_method and option_method.type = 2"
	sql += " left join collateral on auctions.coll_id  = collateral.coll_id"
	sql += " left join useraccount on useraccount.user_id  = auctions.accepted_bidder where to_char(auctions.due_date::date,'yyyy-mm-dd') = to_char('"+api.GetStrings("due_date")[0]+"'::date,'yyyy-mm-dd')'"

	fmt.Println(sql)
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Auctions
	}
	api.ServeJSON()
}

func AllAuctionsCheck(api *AuctionsController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	ul := &getAuctions{}
	json.Unmarshal(frm, ul)
	fmt.Println("sdsdsdsd",ul.From_date)

	Coll_id := ul.Coll_id
	From_date := ul.From_date
	Due_date := ul.Due_date
	Auction_method := ul.Auction_method
	Down_payment := ul.Down_payment
	Description := ul.Description
	Auction_count := ul.Auction_count
	Status := ul.Status

	u := &getAuctions{Coll_id, From_date, Due_date, Auction_method, Down_payment, Description, Auction_count, Status}
	valid.Required(u.Coll_id, "Coll_id")
	valid.Required(u.From_date, "From_date")
	valid.Required(u.Due_date, "Due_date")
	valid.Required(u.Auction_method, "Auction_method")
	valid.Required(u.Down_payment, "Down_payment")
	valid.Required(u.Description, "Description")
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

func AllAuctionsCheckWithId(api *AuctionsController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	Auction_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	ul := &getAuctionsWithId{}
	json.Unmarshal(frm, ul)
	
	Coll_id := ul.Coll_id
	From_date := ul.From_date
	Due_date := ul.Due_date
	Auction_method := ul.Auction_method
	Down_payment := ul.Down_payment
	Description := ul.Description
	Auction_count := ul.Auction_count
	Status := ul.Status

	u := &getAuctionsWithId{Auction_id,Coll_id, From_date, Due_date, Auction_method, Down_payment, Description, Auction_count, Status}
	valid.Required(u.Auction_id, "Auction_id")
	valid.Required(u.Coll_id, "Coll_id")
	valid.Required(u.From_date, "From_date")
	valid.Required(u.Due_date, "Due_date")
	valid.Required(u.Auction_method, "Auction_method")
	valid.Required(u.Down_payment, "Down_payment")
	valid.Required(u.Description, "Description")
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

func (api *AuctionsController) CreateAuctions() {
	frm := api.Ctx.Input.RequestBody

	if AllAuctionsCheck(api) != "" {
		api.Data["json"] = AllAuctionsCheck(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	u := &getAuctions{}
	json.Unmarshal(frm, u)
	// Auction_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	fmt.Println(u)
	Coll_id := u.Coll_id
	From_date := u.From_date
	Due_date := u.Due_date
	Auction_method := u.Auction_method
	Down_payment := u.Down_payment
	Description := u.Description
	Auction_count := u.Auction_count
	Status := u.Status
	From_datenn, _ := time.Parse("2006-01-02", From_date)
	Due_datenn, _ := time.Parse("2006-01-02", Due_date)
	PostsQry := models.Auctions{Coll_id: Coll_id, From_date: From_datenn, Due_date: Due_datenn, Auction_method: Auction_method, Down_payment: Down_payment, Description: Description, Auction_count: Auction_count, Status: Status}
	_, err := o.Insert(&PostsQry)
	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully save data "
	api.ServeJSON()
}

func (api *AuctionsController) EditAuctions() {
    frm := api.Ctx.Input.RequestBody

	if AllAuctionsCheckWithId(api) != "" {
		api.Data["json"] = AllAuctionsCheckWithId(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	u := &getAuctions{}
	json.Unmarshal(frm, u)
	// Auction_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	Auction_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	Coll_id := u.Coll_id
	From_date := u.From_date
	Due_date := u.Due_date
	Auction_method := u.Auction_method
	Down_payment := u.Down_payment
	Description := u.Description
	Auction_count := u.Auction_count
	Status := u.Status
	From_datenn, _ := time.Parse("2006-01-02", From_date)
	Due_datenn, _ := time.Parse("2006-01-02", Due_date)
	PostsQry := models.Auctions{Auction_id: Auction_id,Coll_id: Coll_id, From_date: From_datenn, Due_date: Due_datenn, Auction_method: Auction_method, Down_payment: Down_payment, Description: Description, Auction_count: Auction_count, Status: Status}
	_, err := o.Update(&PostsQry)
	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully update data "
	api.ServeJSON()
}

func (api *AuctionsController) DeleteAuction() {
	chk, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
    if chk==0 {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Auctions [] ambilAuctions
	var sql string
	sql = "select auction_id from auctions where auction_id = "+api.Ctx.Input.Param(":id")+"";
	fmt.Println(sql)
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err == orm.ErrNoRows ||  num == 0 {
		api.Data["json"] = "not found"
	    api.ServeJSON()   
	}
	sql = "UPDATE auctions SET status = 2 where auction_id = "+api.Ctx.Input.Param(":id")	
	o.Raw(sql).QueryRows(&Auctions)
	api.Data["json"] = "successfully delete auctions with auction_id = "+api.Ctx.Input.Param(":id")
	api.ServeJSON()
}

func (api *AuctionsController) ExtendDueDate() {
    chk := api.GetString("due_date")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
	auction_id := api.GetString("auction_id")
    if auction_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Auctions []*models.Auctions
	var sql string
	sql = "select auctions_id from auctions where auctions_id = '"+api.GetString("auction_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE auctions SET auction_count = auction_count+1,due_date = '"+api.GetStrings("due_date")[0]+"' where auction_id = '"+api.GetStrings("auction_id")[0]+"'"
	o.Raw(sql).QueryRows(&Auctions)
	api.Data["json"] = "successfully extend date with auction_id = "+api.GetStrings("auction_id")[0]
	api.ServeJSON()
}

func (api *AuctionsController) ChooseBid() {
    accepted_bidder := api.GetString("accepted_bidder")
    if accepted_bidder == "" {
        api.Ctx.WriteString("")
        return
    }
	auction_id := api.GetString("auction_id")
    if auction_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Auctions []*models.Auctions
	var sql string
	sql = "select auctions_id from auctions where auctions_id = '"+api.GetString("auction_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE auctions SET accepted_bidder = '"+api.GetStrings("accepted_bidder")[0]+"' where auction_id = '"+api.GetStrings("auction_id")[0]+"'"
	o.Raw(sql).QueryRows(&Auctions)
	api.Data["json"] = "successfully accept bidder to accepted_bidder = "+api.GetStrings("accepted_bidder")[0]+" on auction_id = "+api.GetStrings("auction_id")[0]
	api.ServeJSON()
}

func (api *AuctionsController) ChangeAuctionToProcess() {
    chk := api.GetString("auction_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Auctions []*models.Auctions
	var sql string
	sql = "select auctions_id from auctions where auctions_id = '"+api.GetString("auction_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Auctions)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE auctions SET Status  = 2 where auction_id = '"+api.GetStrings("auction_id")[0]+"'"
	o.Raw(sql).QueryRows(&Auctions)
	api.Data["json"] = "successfully change auction to process with auction_id = "+api.GetStrings("auction_id")[0]
	api.ServeJSON()
}