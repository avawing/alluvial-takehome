package main

import (
	"alluvial/handlers"
	"alluvial/repository"
	"alluvial/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"time"

	"fmt"
	"log"
	"net/http"
)

var INFURA_API_KEY string
var ALCHEMY_API_KEY string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	INFURA_API_KEY = os.Getenv("INFURA_API_KEY")
	ALCHEMY_API_KEY = os.Getenv("ALCHEMY_API_KEY")
}

func inject() *gin.Engine {
	router := gin.Default()

	S := services.NewInfuraService(&services.ClientConfig{
		InfuraRepository: repository.NewInfuraRepository(
			&http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       10 * time.Second},
			INFURA_API_KEY,
		)})
	A := services.NewAlchemyService(&services.ClientConfig{
		AlchemyRepository: repository.NewAlchemyRepository(
			&http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       10 * time.Second,
			},
			ALCHEMY_API_KEY),
	})

	handlers.NewHandler(
		&handlers.Config{
			R:              router,
			InfuraService:  S,
			AlchemyService: A,
		})
	return router
}

func main() {
	router := inject()

	router.Run(":8080")
	fmt.Println("Server listening on port 8080")
}
