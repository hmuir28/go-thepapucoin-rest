package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hmuir28/go-thepapucoin-rest/controller"
)

func TransferRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/transfer", controller.SendMoney())
}
