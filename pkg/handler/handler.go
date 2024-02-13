package handler

import (
	"github.com/MrSaveliy/store/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serveces *service.Service) *Handler {
	return &Handler{services: serveces}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	profiles := router.Group("/profiles", h.userIdentity)
	{
		profiles.POST("/", h.createProfile)
		profiles.GET("/:id", h.getProfileById)
		profiles.PUT("/:id", h.updateProfile)
		profiles.DELETE("/:id", h.deleteProfile)
	}
	products := router.Group("/product")
	{
		products.POST("/", h.createProduct)
		products.GET("/:id", h.getProductById)
		products.PUT("/:id", h.updateProduct)
		products.DELETE("/:id", h.deleteProduct)
	}
	categorys := router.Group("/category")
	{
		categorys.POST("/", h.createCategory)
		categorys.GET("/", h.getCategories)
		categorys.GET("/:id", h.getCategoryById)
		categorys.PUT("/:id", h.updateCategory)
		categorys.DELETE("/:id", h.deleteCategory)
	}
	orders := router.Group("/order", h.userIdentity)
	{
		orders.POST("/", h.createOrder)
		orders.GET("/:id", h.getOrderById)
		orders.PUT("/:id", h.updateOrder)
		orders.DELETE("/:id", h.deleteOrder)
	}
	bascket := router.Group("/bascket")
	{
		bascket.POST("/", h.addProductToBascket)
		bascket.GET("/:id", h.getBascketByUserId)
		bascket.DELETE("/:id", h.deleteProductFromBascket)
	}

	favorites := router.Group("/favorites")
	{
		favorites.POST("/", h.addProductToFavorites)
		favorites.GET("/:id", h.getFavoritesByUserId)
		favorites.DELETE("/:id", h.deleteProductFromFavorites)
	}
	return router
}
