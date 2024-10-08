package main

import (
	"fmt"
	"go-shop/db/supabase"
	"go-shop/internal/repository"
	"go-shop/internal/router"
	"go-shop/pkg/config"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("cannot load config", err)
	}
	ApiUrl := "https://pyxekqjrxumchztosmno.supabase.co"
	ApiKey := config.GetEnv("SUPABASE_KEY", "SUPABASE_KEY")
	db, err := supabase.SetupDB(ApiUrl, ApiKey)
	repo := repository.NewSupabaseRepository(db)

	if err != nil {
		fmt.Println("cannot setup supabase client", err)
	}
	r := router.NewRouter(e, repo)
	r.SetupRouter()

	err = e.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
}
