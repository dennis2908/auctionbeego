package routers

import (
	"github.com/astaxie/beego"
	"api_beego/controllers"
)

func init() {

	beego.Router("/GetAllCollateral", &controllers.CollateralController{}, "get:GetAllCollateral")
	beego.Router("/GetCollateralByID", &controllers.CollateralController{}, "get:GetCollateralByID")
	beego.Router("/GetCollateralByCategory", &controllers.CollateralController{}, "get:GetCollateralByCategory")
	beego.Router("/GetCollateralByLJK", &controllers.CollateralController{}, "get:GetCollateralByLJK")
	beego.Router("/GetCollateralByLocation", &controllers.CollateralController{}, "get:GetCollateralByLocation")
	beego.Router("/CreateCollateral", &controllers.CollateralController{}, "post:CreateCollateral")
	beego.Router("/EditCollateral/:id", &controllers.CollateralController{}, "put:EditCollateral")
	
}