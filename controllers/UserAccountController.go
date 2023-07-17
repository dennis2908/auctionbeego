package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "api_beego/models"
	"github.com/badoux/checkmail"
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
	Status_name string
}

func (api *UserAccountController) CreateUser() {
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
	ktp_no := api.GetString("ktp_no")
    if ktp_no == "" {
        api.Ctx.WriteString("")
        return
    }
	status := api.GetString("status")
    if status == "" {
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
	var User [] ambilUser
	var sql string
	sql = "select * from useraccount"
	sql += " where username = '"+api.GetStrings("username")[0]+"' and email = '"+api.GetStrings("email")[0]+"'"
	num, err := o.Raw(sql).QueryRows(&User)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "INSERT INTO useraccount (username, password,email, status, ktp_no) VALUES ('"+api.GetStrings("username")[0]+"'"
	sql +=", '"+api.GetStrings("password")[0]+"','"+api.GetStrings("email")[0]+"', '"+api.GetStrings("status")[0]+"', '"+api.GetStrings("ktp_no")[0]+"')" 	
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = 1
	api.ServeJSON()
}

func (api *UserAccountController) EditUser() {
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
	ktp_no := api.GetString("ktp_no")
    if ktp_no == "" {
        api.Ctx.WriteString("")
        return
    }
	status := api.GetString("status")
    if status == "" {
        api.Ctx.WriteString("")
        return
    }
	email := api.GetString("email")
    if email == "" {
        api.Ctx.WriteString("")
        return
    }
	user_id := api.GetString("user_id")
    if user_id == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var User []*models.UserAccount
	var sql string
	sql = "select useraccount from useraccount where user_id = '"+api.GetString("user_id")+"'";
	num, err := o.Raw(sql).QueryRows(&User)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE useraccount SET username = '"+api.GetStrings("username")[0]+"',password = '"+api.GetStrings("password")[0]+"'"
	sql += ",ktp_no = '"+api.GetStrings("ktp_no")[0]+"',status = '"+api.GetStrings("status")[0]+"',email = '"+api.GetStrings("email")[0]+"' where user_id = '"+api.GetStrings("user_id")[0]+"'"
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = "success edit user account with id = "+api.GetStrings("user_id")[0]
	api.ServeJSON()
}

func (api *UserAccountController) DeleteUser() {
    chk := api.GetString("user_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var User []*models.UserAccount
	var sql string
	sql = "select useraccount from useraccount where user_id = '"+api.GetString("user_id")+"'";
	num, err := o.Raw(sql).QueryRows(&User)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE useraccount SET status = 3 where user_id = '"+api.GetStrings("user_id")[0]+"'"	
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = "success delete user account with id = "+api.GetStrings("user_id")[0]		
	api.ServeJSON()
}

func (api *UserAccountController) DeactivateUser() {
    chk := api.GetString("user_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var User []*models.UserAccount
	var sql string
	sql = "select useraccount from useraccount where user_id = '"+api.GetString("user_id")+"'";
	num, err := o.Raw(sql).QueryRows(&User)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE useraccount SET status = 2 where user_id = '"+api.GetStrings("user_id")[0]+"'" 	
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = "success delete user account with id = "+api.GetStrings("user_id")[0]		
	api.ServeJSON()
}

func (api *UserAccountController) AuthenticateUserX() {
    chk := api.GetString("email")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    err := checkmail.ValidateHost(api.GetStrings("email")[0])
	api.Data["json"] = err		
	api.ServeJSON()
}

func (api *UserAccountController) AuthenticateUser() {
    chk := api.GetString("user_id")
    if chk == "" {
        api.Ctx.WriteString("")
        return
    }
    o := orm.NewOrm()
	o.Using("default")
	var User []*models.UserAccount
	var sql string
	sql = "select useraccount from useraccount where user_id = '"+api.GetString("user_id")+"'";
	num, err := o.Raw(sql).QueryRows(&User)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = ""
	    api.ServeJSON()   
	}
	sql = "UPDATE useraccount SET status = 2 where user_id = '"+api.GetStrings("user_id")[0]+"'" 	
	o.Raw(sql).QueryRows(&User)
	api.Data["json"] = "success delete user account with id = "+api.GetStrings("user_id")[0]		
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
	sql = "select name as status_name,useraccount.* from useraccount"
	sql += " join m_option on useraccount.status = m_option.no and m_option.type = 5"
	sql += " where username = '"+api.GetStrings("username")[0]+"' and password = '"+api.GetStrings("password")[0]+"' and status = 2"
	num, err := o.Raw(sql).QueryRows(&user)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = user[0]
	}
	api.ServeJSON()
}