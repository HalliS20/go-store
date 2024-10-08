package repository

import (
	"encoding/json"
	"go-shop/internal/domain/interfaces"
	"go-shop/internal/models"
	"go-shop/pkg/config"
	"log"
	"strconv"

	"github.com/supabase-community/supabase-go"
)

type SupabaseRepository struct {
	db *supabase.Client
}

func NewSupabaseRepository(database *supabase.Client) interfaces.IProductRepository {
	return &SupabaseRepository{db: database}
}

func (s *SupabaseRepository) GetProducts() []models.Product {
	data, _, err := s.db.From("products").Select("*", "exact", false).Execute()
	if err != nil {
		log.Println("Error getting products:", err)
	}
	log.Println("Data:", string(data))
	var products []models.Product
	if err := json.Unmarshal(data, &products); err != nil {
		log.Println("Error unmarshaling data:", err)
	}
	var fixedProducts []models.Product
	url := config.GetEnv("PHOTO_BUCKET_URL", "PHOTO_BUCKET_URL")
	prodLen := len(products)
	for i := 0; i < prodLen; i++ {
		prod := products[prodLen-i-1]
		prod.Photo = url + prod.Photo
		fixedProducts = append(fixedProducts, prod)
	}

	return fixedProducts
}

func (s *SupabaseRepository) GetProduct(id int) models.Product {
	data, _, err := s.db.From("products").Select("*", "exact", false).Eq("id", strconv.Itoa(id)).Single().Execute()
	if err != nil {
		log.Println("Error getting product:", err)
	}
	var product models.Product
	if err := json.Unmarshal(data, &product); err != nil {
		log.Println("Error unmarshaling data:", err)
	}
	return product
}
