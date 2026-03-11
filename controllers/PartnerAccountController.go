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
	"github.com/badoux/checkmail"
)

type PartnerAccountController struct {
	beego.Controller
}

type ambilPartner struct {
    Partner_id int
	Username string
	Password string
	Email string
	Ljk_id int	
	Status int
	Ktp_no string
	Status_name string
}

func (api *PartnerAccountController) CreatePartner() {
    username := api.GetString("username")
    if username == "" {
        api.Ctx.WriteString("")
        return
    }
	ljk_id := api.GetString("ljk_id")
    if ljk_id == "" {
        api.Ctx.WriteString("")
        return
    }
	password := api.GetString("password")
    if password == "" {
        api.Ctx.WriteString("")
        return
    }
	status := api.GetString("status")
    if status == "" {
        api.Ctx.WriteString("")
        return
    }
	ktp_no := api.GetString("ktp_no")
    if ktp_no == "" {
        api.Ctx.WriteString("")
        return
    }
	email := api.GetString("email")
    if email == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Partner [] ambilPartner
	var sql string
	sql = "select * from partneraccount"
	sql += " where username = '"+api.GetStrings("username")[0]+"' or email = '"+api.GetStrings("email")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Partner)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "INSERT INTO partneraccount (username, ljk_id, password, status, ktp_no, email) VALUES ('"+api.GetStrings("username")[0]+"', '"
	sql += api.GetStrings("ljk_id")[0]+"','"+api.GetStrings("password")[0]+"', '"+api.GetStrings("status")[0]+"', '"+api.GetStrings("ktp_no")[0]+"', '"+api.GetStrings("email")[0]+"')" 	
	o.Raw(sql).QueryRows(&Partner)
	api.Data["json"] = 1
	api.ServeJSON()
	
}

func (api *PartnerAccountController) EditPartner() {
    username := api.GetString("username")
    if username == "" {
        api.Ctx.WriteString("")
        return
    }
	ljk_id := api.GetString("ljk_id")
    if ljk_id == "" {
        api.Ctx.WriteString("")
        return
    }
	password := api.GetString("password")
    if password == "" {
        api.Ctx.WriteString("")
        return
    }
	ktp_no := api.GetString("ktp_no")
    if ktp_no == "" {
        api.Ctx.WriteString("")
        return
    }
	partner_id := api.GetString("partner_id")
    if partner_id == "" {
        api.Ctx.WriteString("")
        return
    }
	email := api.GetString("email")
    if email == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Partner []*models.PartnerAccount
	var sql string
	sql = "select partner_id from partneraccount where partner_id = '"+api.GetString("partner_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Partner)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE partneraccount SET username = '"+api.GetStrings("username")[0]+"',password = '"+api.GetStrings("password")[0]
	sql += "',ljk_id = '"+api.GetStrings("ljk_id")[0]+"',ktp_no = '"+api.GetStrings("ktp_no")[0]
	sql += "',ljk_id = '"+api.GetStrings("ljk_id")[0]+"',email = '"+api.GetStrings("email")[0]+"' where partner_id = '"+api.GetStrings("partner_id")[0]+"'"	
	o.Raw(sql).QueryRows(&Partner)
	api.Data["json"] = "successfully edit partner account with id = "+api.GetStrings("partner_id")[0]
	api.ServeJSON()
}

func (api *PartnerAccountController) DeletePartner() {
    chk := api.GetString("partner_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Partner []*models.PartnerAccount
	var sql string
	sql = "select partner_id from partneraccount where partner_id = '"+api.GetString("partner_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Partner)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE partneraccount SET status = 3 where partner_id = '"+api.GetStrings("partner_id")[0]+"'"	
	o.Raw(sql).QueryRows(&Partner)
	api.Data["json"] = "success delete partner account with id = "+api.GetStrings("partner_id")[0]		
	api.ServeJSON()
}

func (api *PartnerAccountController) DeactivatePartner() {
    chk := api.GetString("partner_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Partner []*models.PartnerAccount
	var sql string
	sql = "select partner_id from partneraccount where partner_id = '"+api.GetString("partner_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Partner)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE partneraccount SET status = 2 where partner_id = '"+api.GetStrings("partner_id")[0]+"'" 	
	o.Raw(sql).QueryRows(&Partner)
	api.Data["json"] = "success deactivated partner account with id = "+api.GetStrings("partner_id")[0]		
	api.ServeJSON()
}

func (api *PartnerAccountController) PartnerLogin() {
    username := api.GetString("username")
    if username == "" {
        api.Ctx.WriteString("")
        return
    }
	password := api.GetString("password")
    if password == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Partner[] ambilPartner
	var sql string
	sql = "select name as status_name,partneraccount.* from partneraccount"
	sql += " join m_option on partneraccount.status = m_option.no and m_option.type = 5"
	sql += " where username = '"+api.GetStrings("username")[0]+"' and password = '"+api.GetStrings("password")[0]+"' and status = 2"
	num, err := o.Raw(sql).QueryRows(&Partner)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Partner[0]
	}
	api.ServeJSON()
}

func (api *PartnerAccountController) AuthenticatePartnerX() {
    chk := api.GetString("email")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    err := checkmail.ValidateHost(api.GetStrings("email")[0])
	api.Data["json"] = err		
	api.ServeJSON()
}

func (api *PartnerAccountController) AuthenticatePartner() {
    chk := api.GetString("partner_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Partner []*models.PartnerAccount
	var sql string
	sql = "select partner_id from partneraccount where partner_id = '"+api.GetString("partner_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Partner)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE partneraccount SET status = 2 where partner_id = '"+api.GetStrings("partner_id")[0]+"'" 	
	o.Raw(sql).QueryRows(&Partner)
	api.Data["json"] = "success deactivated partner account with id = "+api.GetStrings("partner_id")[0]		
	api.ServeJSON()
}