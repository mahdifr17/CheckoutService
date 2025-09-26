package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mahdifr17/CheckoutService/internals/config"
	"github.com/mahdifr17/CheckoutService/internals/handlers"
	"github.com/mahdifr17/CheckoutService/internals/infra/repository"
	"github.com/mahdifr17/CheckoutService/internals/services"
	"gorm.io/gorm"
)

func StartAPI(cfg *config.Config, db *gorm.DB) {
	// Initialize repositories
	productRepo := repository.NewProductRepository(db)
	promotionRepo := repository.NewPromotionRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	// Initialize services
	checkoutService := services.NewCheckoutService(productRepo, promotionRepo, orderRepo)

	// Initialize handlers
	checkoutHandler := handlers.NewCheckoutHandler(checkoutService)

	// Setup router
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API routes
	api := r.Group("/v1")
	{
		api.GET("/products", checkoutHandler.GetProducts)
		api.POST("/checkout", checkoutHandler.Checkout)
	}

	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Fatal(r.Run(":" + cfg.Server.Port))
}
