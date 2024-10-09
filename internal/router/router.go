package router

import (
	"go-shop/internal/domain/interfaces"
	"log"
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

		c.HTML(http.StatusOK, "index.html", gin.H{"title": "SHOE", "products": products, "favicon": favicon})
	})

	r.e.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		product := products[idInt-1]
		c.HTML(http.StatusOK, "product.html", gin.H{"title": "SHOE", "product": product, "favicon": favicon})
	})

	r.e.GET("/nav", func(c *gin.Context) {
		// Read the template file
		data := &gin.H{"products": products, "favicon": favicon}
		prodList, err := reloadHtml("templates/product_list.html", data)
		if err != nil {
			log.Printf("Error reloading HTML: %v", err)
			c.String(http.StatusInternalServerError, "Error reloading HTML")
			return
		}
		// Log the rendered HTML
		log.Printf("Rendered HTML for /nav:\n%s", prodList)
		returnJson := gin.H{"html": prodList, "title": "SHOE", "description": "SHOE.com - The best place to buy shoes online"}

		c.JSON(http.StatusOK, returnJson)
	})

	r.e.GET("/nav/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		product := products[idInt-1]

		data := &gin.H{"product": product, "favicon": favicon}
		prodDetails, err := reloadHtml("templates/product_details.html", data)
		if err != nil {
			log.Printf("Error reloading HTML: %v", err)
			c.String(http.StatusInternalServerError, "Error reloading HTML")
			return
		}
		log.Printf("Rendered HTML for /nav/products/:id :\n%s", prodDetails)
		returnJson := gin.H{"html": prodDetails, "description": product.Description, "title": "SHOE | " + product.Name}

		c.JSON(http.StatusOK, returnJson)
	})

	r.e.GET("/public/*filepath", func(c *gin.Context) {
		c.File("public/" + c.Param("filepath"))
	})

	r.e.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

}
