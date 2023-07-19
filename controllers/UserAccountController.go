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
	_ "github.com/badoux/checkmail"
)

type UserAccountController struct {
	beego.Controller
}

type ambilUser struct {
    User_id int
	Username string
	Password string	
	Email string
	Status int
	Ktp_no string
}

type cekUser struct {
  	Username string
	Password string	
	Email string
	Status int
	Ktp_no string
}

type email struct {
  Email string
}

func AllUserCheck(api *UserAccountController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	ul := &cekUser{}
	json.Unmarshal(frm, ul)

	Username := ul.Username
	Password := ul.Password
	Email := ul.Email
	Status := ul.Status
	Ktp_no := ul.Ktp_no

	u := &cekUser{Username, Password, Email, Status, Ktp_no}
	valid.Required(u.Username, "Username")
	valid.Required(u.Password, "Password")
	valid.Required(u.Email, "Email")
	valid.Required(u.Status, "Status")
	valid.Required(u.Ktp_no, "Ktp_no")

	if valid.HasErrors() {
		// If there are error messages it means the validation didn't pass
		// Print error message
		for _, err := range valid.Errors {
			return err.Key + err.Message
		}
	}

	return ""
}


func (api *UserAccountController) CreateUser() {
	frm := api.Ctx.Input.RequestBody

	if AllUserCheck(api) != "" {
		api.Data["json"] = AllUserCheck(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	u:= &cekUser{}
	json.Unmarshal(frm,u)
	Username := u.Username
	Password := u.Password
	Email := u.Email
	Status := u.Status
	Ktp_no := u.Ktp_no
	PostsQry := models.UserAccount{Username: Username, Password: Password, Email: Email, Status: Status, Ktp_no: Ktp_no}
	_, err := o.Insert(&PostsQry)
	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully save data"
	api.ServeJSON()
}

func AllUserCheckWithId(api *UserAccountController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	User_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	ul := &ambilUser{}
	json.Unmarshal(frm, ul)
	Username := ul.Username
	Password := ul.Password
	Email := ul.Email
	Status := ul.Status
	Ktp_no := ul.Ktp_no

	u := &ambilUser{User_id,Username,Password,Email, Status, Ktp_no}
	valid.Required(u.User_id, "User_id")
	valid.Required(u.Username, "Username")
	valid.Required(u.Password, "Password")
	valid.Required(u.Email, "Email")
	valid.Required(u.Status, "Status")
	valid.Required(u.Ktp_no, "Ktp_no")

	if valid.HasErrors() {
		// If there are error messages it means the validation didn't pass
		// Print error message
		for _, err := range valid.Errors {
			return err.Key + err.Message
		}
	}

	return ""
}

func (api *UserAccountController) EditUser() {
    frm := api.Ctx.Input.RequestBody

	if AllUserCheckWithId(api) != "" {
		api.Data["json"] = AllUserCheckWithId(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	u := &ambilUser{}
	json.Unmarshal(frm, u)
	User_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	Username := u.Username
	Password := u.Password
	Email := u.Email
	Status := u.Status
	Ktp_no := u.Ktp_no
	PostsQry := models.UserAccount{User_id: User_id,Username: Username, Password: Password, Email: Email, Status: Status, Ktp_no: Ktp_no}
	_, err := o.Update(&PostsQry)
	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully update data"
	api.ServeJSON()
}

func (api *UserAccountController) DeleteUser() {
    if api.Ctx.Input.Param(":id") == "" {
        api.Ctx.WriteString("Id is needed")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var User []*models.UserAccount
	var sql string
	sql = "select useraccount from useraccount where user_id = '"+api.Ctx.Input.Param(":id")+"'";
	num, err := o.Raw(sql).QueryRows(&User)
	if err == orm.ErrNoRows && num == 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE useraccount SET status = 3 where user_id = '"+api.Ctx.Input.Param(":id")+"'"	
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = "success delete user account with id = '"+api.Ctx.Input.Param(":id")+"'"			
	api.ServeJSON()
}

func (api *UserAccountController) DeactivateUser() {
    // User_id, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
    o := orm.NewOrm()
	o.Using("default")
	var User []*models.UserAccount
	var sql string
	sql = "select useraccount from useraccount where user_id = "+api.Ctx.Input.Param(":id")
	num, err := o.Raw(sql).QueryRows(&User)
	if err == orm.ErrNoRows || num == 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE useraccount SET status = 2 where user_id = "+api.Ctx.Input.Param(":id")
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = "success delete user account with id = "+api.Ctx.Input.Param(":id")	
	api.ServeJSON()
}

// func (api *UserAccountController) AuthenticateUserX() {
//     chk := api.GetString("email")

// 	frm := api.Ctx.Input.RequestBody
// 	ul := &email{}
// 	json.Unmarshal(frm, ul)
//     err := checkmail.ValidateHost(api.GetStrings("email")[0])
// 	api.Data["json"] = err		
// 	api.ServeJSON()
// }

func (api *UserAccountController) AuthenticateUser() {
	frm := api.Ctx.Input.RequestBody
	ul := &email{}
	json.Unmarshal(frm, ul)
    if ul.Email == "" {
        api.Ctx.WriteString("Email is empty")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var User []*models.UserAccount
	var sql string
	sql = "select useraccount from useraccount where user_id = '"+ul.Email+"'";
	num, err := o.Raw(sql).QueryRows(&User)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE useraccount SET status = 2 where user_id = '"+ul.Email+"'" 	
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = "success delete user account with id = '"+ul.Email+"'"	
	api.ServeJSON()
}

func (api *UserAccountController) UserLogin() {

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
	var user[ ] ambilUser
	var sql string
	sql = "select useraccount.* from useraccount"
	sql += " join m_option on useraccount.status = m_option.no and m_option.type = 5"
	sql += " where username = '"+api.GetStrings("username")[0]+"' and password = '"+api.GetStrings("password")[0]+"' and status = 2"
	num, err := o.Raw(sql).QueryRows(&user)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = user[0]
	}
	api.ServeJSON()
}