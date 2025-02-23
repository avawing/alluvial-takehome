package main

import (
	"alluvial/handlers"
	"alluvial/repository"
	"alluvial/services"
	"alluvial/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var INFURA_API_KEY string
var ALCHEMY_API_KEY string
var CHAINSTACK_API_KEY string
var CHAINSTACK_API_NODE string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	INFURA_API_KEY = os.Getenv("INFURA_API_KEY")
	ALCHEMY_API_KEY = os.Getenv("ALCHEMY_API_KEY")
	CHAINSTACK_API_KEY = os.Getenv("CHAINSTACK_API_KEY")
	CHAINSTACK_API_NODE = os.Getenv("CHAINSTACK_NODE")
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

	C := services.NewChainstackService(&services.ClientConfig{
		ChainstackRepository: repository.NewChainstackRepository(
			&http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       10 * time.Second,
			},
			CHAINSTACK_API_KEY,
			CHAINSTACK_API_NODE,
		),
	})

	lb := utils.NewLoadBalancer([]utils.Server{A, S, C})

	handlers.NewHandler(
		&handlers.Config{
			R:                 router,
			InfuraService:     S,
			AlchemyService:    A,
			ChainstackService: C,
			LoadBalancer:      lb,
		})
	return router
}

func main() {
	router := inject()

	router.Run(":8080")
	fmt.Println("Server listening on port 8080")
}
