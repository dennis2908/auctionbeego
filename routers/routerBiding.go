package routers

import (
	"github.com/astaxie/beego"
	"api_beego/controllers"
)

func init() {
	beego.Router("/GetAllBids", &controllers.BiddingsController{}, "post:GetAllBids")
	beego.Router("/GetBidsByAuction", &controllers.BiddingsController{}, "get:GetBidsByAuction")
	beego.Router("/CreateBiddings", &controllers.BiddingsController{}, "post:CreateBiddings")
	beego.Router("/CancelBid/id", &controllers.BiddingsController{}, "put:CancelBid")
	beego.Router("/PayBidDownPayment/:id", &controllers.BiddingsController{}, "put:PayBidDownPayment")
}