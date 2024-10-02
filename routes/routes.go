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

	// Invoice routes
	r.POST("/invoices", controllers.CreateInvoice)
	r.GET("/invoices/:invoiceId", controllers.CheckInvoiceStatus)
	r.PUT("/invoices/:invoiceId", controllers.UpdateInvoice)
	r.DELETE("/invoices/:invoiceId", controllers.DeleteInvoice)

	// Subscription routes
	r.POST("/subscriptions", controllers.CreateSubscription)
	r.GET("/subscriptions/:planId/:customerId", controllers.GetSubscription) // Adjusted route
	r.PUT("/subscriptions/:planId/:customerId", controllers.UpdateSubscription) // Adjusted route
	r.DELETE("/subscriptions/:planId/:customerId", controllers.DeleteSubscription) // Adjusted route

	// Login route
	r.POST("/auth/login", controllers.Login)
}
