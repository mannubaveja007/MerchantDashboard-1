package routes

import (
	"merchant-dashboard/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Product routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts) // Adjusted to get products by merchant_id as a query parameter
	r.PUT("/products/:merchantId/:productId", controllers.UpdateProduct)
	r.DELETE("/products/:merchantId/:productId", controllers.DeleteProduct)

	// Payment routes
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
}
