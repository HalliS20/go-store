package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Router) getProducts(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=172800")
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "SHOE", "products": r.products, "favicon": r.favicon})
}

func (r *Router) getProduct(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=172800")
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	product := r.products[idInt-1]
	c.HTML(http.StatusOK, "product.html", gin.H{"title": "SHOE", "product": product, "favicon": r.favicon})

}

func (r *Router) navToHome(c *gin.Context) {
	// Read the template file
	data := &gin.H{"products": r.products}
	prodList, err := reloadHtml("templates/product_list.html", data)
	if err != nil {
		log.Printf("Error reloading HTML: %v", err)
		c.String(http.StatusInternalServerError, "Error reloading HTML")
		return
	}
	returnJson := gin.H{"html": prodList, "title": "SHOE", "description": "SHOE.com - The best place to buy shoes online"}

	c.JSON(http.StatusOK, returnJson)
}

func (r *Router) navToProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	product := r.products[idInt-1]

	data := &gin.H{"product": product, "favicon": r.favicon}
	prodDetails, err := reloadHtml("templates/product_details.html", data)
	if err != nil {
		log.Printf("Error reloading HTML: %v", err)
		c.String(http.StatusInternalServerError, "Error reloading HTML")
		return
	}
	returnJson := gin.H{"html": prodDetails, "description": product.Description, "title": "SHOE | " + product.Name}

	c.JSON(http.StatusOK, returnJson)
}
