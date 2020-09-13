package account

import "github.com/gin-gonic/gin"

func (h *Handler) initRoute(r gin.IRouter) {
	v1 := r.Group("v1")
	v1.Use()
	{
		v1.GET("accounts", h.GetAccounts)
		acc := v1.Group("account")
		acc.POST("", h.SaveAccount)
		acc.GET(":id", h.GetAccount)
		acc.PUT(":id", h.UpdateAccount)
		acc.DELETE(":id", h.DeleteAccount)
	}
}
