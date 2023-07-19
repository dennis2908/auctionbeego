package routers

import (
	"github.com/astaxie/beego"
	"api_beego/controllers"
)

func init() {
 	beego.Router("/CreateUser", &controllers.UserAccountController{}, "post:CreateUser")
	beego.Router("/EditUser/:id", &controllers.UserAccountController{}, "put:EditUser")
	beego.Router("/DeleteUser/:id", &controllers.UserAccountController{}, "delete:DeleteUser")
	beego.Router("/DeactivateUser/:id", &controllers.UserAccountController{}, "get:DeactivateUser")
	beego.Router("/UserLogin", &controllers.UserAccountController{}, "post:UserLogin")
	beego.Router("/AuthenticateUser", &controllers.UserAccountController{}, "post:AuthenticateUser")
}