package main

import (
	"alluvial/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func inject() *gin.Engine {
	router := gin.Default()

	handlers.NewHandler(
		&handlers.Config{
			R: router,
		})
	return router
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go app!")
}

func main() {
	router := inject()
	router.Run(":8080")
	fmt.Println("Server listening on port 8080")
}
