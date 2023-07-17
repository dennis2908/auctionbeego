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
	Username string
	Email string
	Ktp_no string
	Type_name string
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

func (api *CollateralController) CreateCollateral() {

	// frm := api.Ctx.Input.RequestBody
	// if AllPostsCheck(api) != "" {
	// 	api.Data["json"] = AllPostsCheck(api)
	// 	api.ServeJSON()
	// 	return
	// }

	// o := orm.NewOrm()
	// o.Using("default")

	// u := &ambilPosts{}
	// json.Unmarshal(frm, u)
	// idInt, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	// Title := u.Title
	// Content := u.Content
	// Category := u.Category
	// Status := u.Status
	// PostsQry := models.Posts{Id: idInt, Title: Title, Content: Content, Category: Category, Status: Status}
    acc_id := api.GetString("acc_id")
    if acc_id == "" {
        api.Ctx.WriteString("")
        return
    } 
	type_id := api.GetString("type_id")
    if type_id == "" {
        api.Ctx.WriteString("")
        return
    } 
	owner_name := api.GetString("owner_name")
    if owner_name == "" {
        api.Ctx.WriteString("")
        return
    } 
	coll_location := api.GetString("coll_location")
    if coll_location == "" {
        api.Ctx.WriteString("")
        return
    } 
	initial_col_price := api.GetString("initial_col_price")
    if initial_col_price == "" {
        api.Ctx.WriteString("")
        return
    } 
	final_col_price := api.GetString("final_col_price")
    if final_col_price == "" {
        api.Ctx.WriteString("")
        return
    } 
	ljk_id := api.GetString("ljk_id")
    if ljk_id == "" {
        api.Ctx.WriteString("")
        return
    }
	document_path := api.GetString("document_path")
    if document_path == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Collateral []*models.Collateral
	var sql string
	sql = "INSERT INTO collateral (acc_id, type_id, owner_name, coll_location, initial_col_price, final_col_price, ljk_id, document_path) VALUES ('"+api.GetStrings("acc_id")[0]+"'"
	sql += ",'"+api.GetStrings("type_id")[0]+"','"+api.GetStrings("owner_name")[0]+"','"+api.GetStrings("coll_location")[0]+"','"+api.GetStrings("initial_col_price")[0]+"'"
	sql += ",'"+api.GetStrings("final_col_price")[0]+"','"+api.GetStrings("ljk_id")[0]+"','"+api.GetStrings("document_path")[0]+"')"
	o.Raw(sql).QueryRows(&Collateral)
	api.Data["json"] = 1
	api.ServeJSON()
}

func (api *CollateralController) EditCollateral() {
    acc_id := api.GetString("acc_id")
    if acc_id == "" {
        api.Ctx.WriteString("")
        return
    }
	type_id := api.GetString("type_id")
    if type_id == "" {
        api.Ctx.WriteString("")
        return
    }
	owner_name := api.GetString("owner_name")
    if owner_name == "" {
        api.Ctx.WriteString("")
        return
    }
	coll_location := api.GetString("coll_location")
    if coll_location == "" {
        api.Ctx.WriteString("")
        return
    }
	initial_col_price := api.GetString("initial_col_price")
    if initial_col_price == "" {
        api.Ctx.WriteString("")
        return
    }
	final_col_price := api.GetString("final_col_price")
    if final_col_price == "" {
        api.Ctx.WriteString("")
        return
    }
	ljk_id := api.GetString("ljk_id")
    if ljk_id == "" {
        api.Ctx.WriteString("")
        return
    }
	document_path := api.GetString("document_path")
    if document_path == "" {
        api.Ctx.WriteString("")
        return
    }
	coll_id := api.GetString("coll_id")
    if coll_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var Collateral []*models.Collateral
	var sql string
	sql = "select coll_id from collateral where coll_id = '"+api.GetString("coll_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE collateral SET acc_id = '"+api.GetStrings("acc_id")[0]+"', type_id = '"+api.GetStrings("type_id")[0]+"', owner_name = '"+api.GetStrings("owner_name")[0]+"'"
	sql += ", coll_location = '"+api.GetStrings("coll_location")[0]+"', initial_col_price = '"+api.GetStrings("initial_col_price")[0]+"', final_col_price = '"+api.GetStrings("final_col_price")[0]+"'"
	sql += ", ljk_id = '"+api.GetStrings("ljk_id")[0]+"', document_path = '"+api.GetStrings("document_path")[0]+"' where coll_id = '"+api.GetStrings("coll_id")[0]+"'" 
	o.Raw(sql).QueryRows(&Collateral)
	api.Data["json"] = "successfully edit collateral with coll_id "+api.GetStrings("coll_id")[0]
	api.ServeJSON()
}