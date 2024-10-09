package router

import (
	"go-shop/internal/domain/interfaces"
	"go-shop/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	e        *gin.Engine
	repo     interfaces.IProductRepository
	products []models.Product
	favicon  string
}

func NewRouter(engine *gin.Engine, repository interfaces.IProductRepository) interfaces.IRouter {
	return &Router{e: engine, repo: repository}
}

func (r *Router) SetupRouter() {
	r.products = r.repo.GetProducts()
	r.favicon = "/public/images/Nail.png"

	r.e.LoadHTMLGlob("templates/*")
	r.e.GET("/", func(c *gin.Context) { r.getProducts(c) })
	r.e.GET("/products/:id", func(c *gin.Context) { r.getProduct(c) })
	r.e.GET("/nav", func(c *gin.Context) { r.navToHome(c) })
	r.e.GET("/nav/products/:id", func(c *gin.Context) { r.navToProduct(c) })
	r.e.GET("/public/*filepath", func(c *gin.Context) {
		c.File("public/" + c.Param("filepath"))
	})
	r.e.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, r.products)
	})

}
