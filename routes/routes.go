package routes

import (
	"merchant-dashboard/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Product routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.PUT("/products/:merchantId/:productId", controllers.UpdateProduct)
	r.DELETE("/products/:merchantId/:productId", controllers.DeleteProduct)

	// Payment routes
	//TODO=> AUTH/JWT
	r.POST("/payments/transfer", controllers.TransferMoney)
	r.POST("/payments/receive", controllers.ReceiveMoney)

	// Invoice routes
	r.POST("/invoices", controllers.CreateInvoice)
	r.GET("/invoices/:invoiceId", controllers.CheckInvoiceStatus)
	r.PUT("/invoices/:invoiceId", controllers.UpdateInvoice)
	r.DELETE("/invoices/:invoiceId", controllers.DeleteInvoice)

	// Subscription routes
	r.POST("/subscriptions", controllers.CreateSubscription)
	r.GET("/subscriptions/:planId", controllers.GetSubscription)
	r.PUT("/subscriptions/:planId", controllers.UpdateSubscription)
	r.DELETE("/subscriptions/:planId", controllers.DeleteSubscription)

	// Login route
	r.POST("/auth/login", controllers.Login)
}
