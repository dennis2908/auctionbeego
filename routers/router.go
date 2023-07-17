package routers

import (
	"github.com/astaxie/beego"
	"api_beego/controllers"
)

func init() {
 	beego.Router("/CreateUser", &controllers.UserAccountController{}, "post:CreateUser")
	beego.Router("/EditUser", &controllers.UserAccountController{}, "post:EditUser")
	beego.Router("/DeleteUser", &controllers.UserAccountController{}, "post:DeleteUser")
	beego.Router("/DeactivateUser", &controllers.UserAccountController{}, "post:DeactivateUser")
	beego.Router("/UserLogin", &controllers.UserAccountController{}, "post:UserLogin")
	beego.Router("/AuthenticateUser", &controllers.UserAccountController{}, "post:AuthenticateUser")
	beego.Router("/CreatePartner", &controllers.PartnerAccountController{}, "post:CreatePartner")
	beego.Router("/EditPartner", &controllers.PartnerAccountController{}, "post:EditPartner")
	beego.Router("/DeletePartner", &controllers.PartnerAccountController{}, "post:DeletePartner")
	beego.Router("/DeactivatePartner", &controllers.PartnerAccountController{}, "post:DeactivatePartner")
	beego.Router("/AuthenticatePartner", &controllers.PartnerAccountController{}, "post:AuthenticatePartner")
	beego.Router("/PartnerLogin", &controllers.PartnerAccountController{}, "post:PartnerLogin")
	beego.Router("/GetAllCollateral", &controllers.CollateralController{}, "post:GetAllCollateral")
	beego.Router("/GetCollateralByID", &controllers.CollateralController{}, "post:GetCollateralByID")
	beego.Router("/GetCollateralByCategory", &controllers.CollateralController{}, "post:GetCollateralByCategory")
	beego.Router("/GetCollateralByLJK", &controllers.CollateralController{}, "post:GetCollateralByLJK")
	beego.Router("/GetCollateralByLocation", &controllers.CollateralController{}, "post:GetCollateralByLocation")
	beego.Router("/CreateCollateral", &controllers.CollateralController{}, "post:CreateCollateral")
	beego.Router("/EditCollateral", &controllers.CollateralController{}, "post:EditCollateral")
	beego.Router("/GetAllBids", &controllers.BiddingsController{}, "post:GetAllBids")
	beego.Router("/GetBidsByAuction", &controllers.BiddingsController{}, "post:GetBidsByAuction")
	beego.Router("/CreateBiddings", &controllers.BiddingsController{}, "post:CreateBiddings")
	beego.Router("/CancelBid", &controllers.BiddingsController{}, "post:CancelBid")
	beego.Router("/PayBidDownPayment", &controllers.BiddingsController{}, "post:PayBidDownPayment")
}