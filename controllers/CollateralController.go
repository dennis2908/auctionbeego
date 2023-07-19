package controllers

import (
	models "api_beego/models"
	_ "context"
	"encoding/json"
	_ "fmt"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/leekchan/accounting"
	_ "github.com/shopspring/decimal"
)

type CollateralController struct {
	beego.Controller
}

type ambilCollateral struct {
    Coll_id int
	Acc_id  int
	Type_id int	
	Owner_name string
	Coll_location string
	Initial_col_price int 
	Final_col_price int
	Ljk_id int
	Document_path string
	User_id int
	Username string
	Password string	
	Email string
	Status int
	Ktp_no string
}

type cekCollateral struct {
    Coll_id int
	Acc_id  int
	Type_id int	
	Owner_name string
	Coll_location string
	Initial_col_price int 
	Final_col_price int
	Ljk_id int
	Document_path string
}

type getCollateral struct {
    Acc_id  int
	Type_id int	
	Owner_name string
	Coll_location string
	Initial_col_price int 
	Final_col_price int
	Ljk_id int
	Document_path string
}

func (api *CollateralController) GetAllCollateral() {
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select collateral.*,useraccount.*,m_option.name as type_name from collateral"
	sql += " left join m_option on m_option.no  = collateral.type_id and type = 1"
	sql += " left join useraccount on useraccount.user_id = collateral.acc_id"
	num, err :=o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Collateral
	}
	api.ServeJSON()
}

func (api *CollateralController) GetCollateralByID() {
    coll_id := api.GetString("coll_id")
    if coll_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select collateral.*,useraccount.*,m_option.name as type_name from collateral"
	sql += " left join useraccount on useraccount.user_id = collateral.acc_id"
	sql += " left join m_option on m_option.no  = collateral.type_id where coll_id = '"+api.GetStrings("coll_id")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Collateral[0]
	}
	api.ServeJSON()
}

func (api *CollateralController) GetCollateralByCategory() {
    type_id := api.GetString("type_id")
    if type_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select collateral.*,useraccount.*,m_option.name as type_name from collateral"
	sql += " left join useraccount on useraccount.user_id = collateral.acc_id"
	sql += " left join m_option on m_option.no  = collateral.type_id where type_id = '"+api.GetStrings("type_id")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Collateral
	}
	api.ServeJSON()
}

func (api *CollateralController) GetCollateralByLJK() {
    ljk_id := api.GetString("ljk_id")
    if ljk_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select collateral.*,useraccount.*,m_option.name as type_name from collateral"
	sql += " left join useraccount on useraccount.user_id = collateral.acc_id"
	sql += " left join m_option on m_option.no  = collateral.type_id where ljk_id = '"+api.GetStrings("ljk_id")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Collateral
	}
	api.ServeJSON()
}

func (api *CollateralController) GetCollateralByLocation() {
    coll_location := api.GetString("coll_location")
    if coll_location == "" {
        api.Ctx.WriteString("")
        return
    } 
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select collateral.*,useraccount.*,m_option.name as type_name from collateral"
	sql += " left join useraccount on useraccount.user_id = collateral.acc_id"
	sql += " left join m_option on m_option.no  = collateral.type_id where coll_location = '"+api.GetStrings("coll_location")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Collateral
	}
	api.ServeJSON()
}

func AllCollateralCheck(api *CollateralController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	ul := &getCollateral{}
	json.Unmarshal(frm, ul)

	Acc_id := ul.Acc_id
	Type_id := ul.Type_id
	Owner_name := ul.Owner_name
	Coll_location := ul.Coll_location
	Initial_col_price := ul.Initial_col_price
	Final_col_price := ul.Final_col_price
	Ljk_id := ul.Ljk_id
	Document_path := ul.Document_path

	u := &getCollateral{Acc_id, Type_id, Owner_name, Coll_location, Initial_col_price, Final_col_price, Ljk_id, Document_path}
	valid.Required(u.Acc_id, "Acc_id")
	valid.Required(u.Type_id, "Type_id")
	valid.Required(u.Owner_name, "Owner_name")
	valid.Required(u.Coll_location, "Coll_location")
	valid.Required(u.Initial_col_price, "Initial_col_price")
	valid.Required(u.Final_col_price, "Final_col_price")
	valid.Required(u.Ljk_id, "Ljk_id")
	valid.Required(u.Document_path, "Document_path")

	if valid.HasErrors() {
		// If there are error messages it means the validation didn't pass
		// Print error message
		for _, err := range valid.Errors {
			return err.Key + err.Message
		}
	}

	return ""
}

func (api *CollateralController) CreateCollateral() {

	frm := api.Ctx.Input.RequestBody

	if AllCollateralCheck(api) != "" {
		api.Data["json"] = AllCollateralCheck(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	u:= &getCollateral{}
	json.Unmarshal(frm,u)
	Acc_id := u.Acc_id
	Type_id := u.Type_id
	Owner_name := u.Owner_name
	Coll_location := u.Coll_location
	Initial_col_price := u.Initial_col_price
	Final_col_price := u.Final_col_price
	Ljk_id := u.Ljk_id
	Document_path := u.Document_path
	PostsQry := models.Collateral{Acc_id: Acc_id, Type_id: Type_id, Owner_name: Owner_name, Coll_location: Coll_location, Initial_col_price: Initial_col_price, Final_col_price: Final_col_price, Ljk_id: Ljk_id, Document_path: Document_path}
	_, err := o.Insert(&PostsQry)
	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully save data "
	api.ServeJSON()
}

func AllCollateralCheckWithId(api *CollateralController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	Coll_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	ul := &getCollateral{}
	json.Unmarshal(frm, ul)
	Acc_id := ul.Acc_id
	Type_id := ul.Type_id
	Owner_name := ul.Owner_name
	Coll_location := ul.Coll_location
	Initial_col_price := ul.Initial_col_price
	Final_col_price := ul.Final_col_price
	Ljk_id := ul.Ljk_id
	Document_path := ul.Document_path

	u := &cekCollateral{Coll_id,Acc_id,Type_id, Owner_name, Coll_location, Initial_col_price, Final_col_price, Ljk_id, Document_path}
	valid.Required(u.Coll_id, "Coll_id")
	valid.Required(u.Acc_id, "Acc_id")
	valid.Required(u.Type_id, "Type_id")
	valid.Required(u.Owner_name, "Owner_name")
	valid.Required(u.Coll_location, "Coll_location")
	valid.Required(u.Initial_col_price, "Initial_col_price")
	valid.Required(u.Final_col_price, "Final_col_price")
	valid.Required(u.Ljk_id, "Ljk_id")
	valid.Required(u.Document_path, "Document_path")

	if valid.HasErrors() {
		// If there are error messages it means the validation didn't pass
		// Print error message
		for _, err := range valid.Errors {
			return err.Key + err.Message
		}
	}

	return ""
}

func (api *CollateralController) EditCollateral() {
	frm := api.Ctx.Input.RequestBody

	if AllCollateralCheckWithId(api) != "" {
		api.Data["json"] = AllCollateralCheckWithId(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	u := &getCollateral{}
	json.Unmarshal(frm, u)
	Coll_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	Acc_id := u.Acc_id
	Type_id := u.Type_id
	Owner_name := u.Owner_name
	Coll_location := u.Coll_location
	Initial_col_price := u.Initial_col_price
	Final_col_price := u.Final_col_price
	Ljk_id := u.Ljk_id
	Document_path := u.Document_path
	PostsQry := models.Collateral{Coll_id: Coll_id,Acc_id: Acc_id, Type_id: Type_id, Owner_name: Owner_name, Coll_location: Coll_location, Initial_col_price: Initial_col_price, Final_col_price: Final_col_price, Ljk_id: Ljk_id, Document_path: Document_path}
	_, err := o.Update(&PostsQry)
	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully update data "
	api.ServeJSON()
}