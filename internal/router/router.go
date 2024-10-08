package router

import (
	"go-shop/internal/domain/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router struct {
	e    *gin.Engine
	repo interfaces.IProductRepository
}

func NewRouter(engine *gin.Engine, repository interfaces.IProductRepository) interfaces.IRouter {
	return &Router{e: engine, repo: repository}
}

func (r *Router) SetupRouter() {
	products := r.repo.GetProducts()
	favicon := "/public/images/Nail.png"

	r.e.LoadHTMLGlob("templates/*")

	r.e.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{"title": "SHOE.com", "products": products, "favicon": favicon})
	})

	r.e.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		product := products[idInt-1]
		c.HTML(http.StatusOK, "product.html", gin.H{"title": "SHOE.com", "product": product, "favicon": favicon})
	})

	r.e.GET("/public/*filepath", func(c *gin.Context) {
		c.File("public/" + c.Param("filepath"))
	})

	r.e.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	r.e.GET("/api/users", func(c *gin.Context) {})

	r.e.GET("/api/users/:id", func(c *gin.Context) {})
}
