package routers

import (
	"github.com/astaxie/beego"
	"api_beego/controllers"
)

func init() {
	beego.Router("/CreatePartner", &controllers.PartnerAccountController{}, "post:CreatePartner")
	beego.Router("/EditPartner", &controllers.PartnerAccountController{}, "post:EditPartner")
	beego.Router("/DeletePartner", &controllers.PartnerAccountController{}, "post:DeletePartner")
	beego.Router("/DeactivatePartner", &controllers.PartnerAccountController{}, "post:DeactivatePartner")
	beego.Router("/AuthenticatePartner", &controllers.PartnerAccountController{}, "post:AuthenticatePartner")
	beego.Router("/PartnerLogin", &controllers.PartnerAccountController{}, "post:PartnerLogin")
	beego.Router("/GetAllBids", &controllers.BiddingsController{}, "post:GetAllBids")
	beego.Router("/GetBidsByAuction", &controllers.BiddingsController{}, "post:GetBidsByAuction")
	beego.Router("/CreateBiddings", &controllers.BiddingsController{}, "post:CreateBiddings")
	beego.Router("/CancelBid", &controllers.BiddingsController{}, "post:CancelBid")
	beego.Router("/PayBidDownPayment", &controllers.BiddingsController{}, "post:PayBidDownPayment")
}