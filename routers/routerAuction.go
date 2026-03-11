package routers

import (
	"github.com/astaxie/beego"
	"api_beego/controllers"
)

func init() {
	beego.Router("/GetAllAuctions", &controllers.AuctionsController{}, "get:GetAllAuctions")
	beego.Router("/GetAuctionsByID", &controllers.AuctionsController{}, "get:GetAuctionsByID")
	beego.Router("/GetAuctionsByCollateral", &controllers.AuctionsController{}, "get:GetAuctionsByCollateral")
	beego.Router("/GetAuctionsByDueDate", &controllers.AuctionsController{}, "get:GetAuctionsByDueDate")
	beego.Router("/DeleteAuction/:id", &controllers.AuctionsController{}, "delete:DeleteAuction")
	beego.Router("/CreateAuctions", &controllers.AuctionsController{}, "post:CreateAuctions")
	beego.Router("/EditAuctions/:id", &controllers.AuctionsController{}, "put:EditAuctions")
	beego.Router("/ExtendDueDate:id", &controllers.AuctionsController{}, "put:ExtendDueDate")
	beego.Router("/ChooseBid:id", &controllers.AuctionsController{}, "put:ChooseBid")
	beego.Router("/ChangeAuctionToProcess:id", &controllers.AuctionsController{}, "put:ChangeAuctionToProcess")
}