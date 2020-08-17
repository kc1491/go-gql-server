package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kc1491/go-gql-server/internal/handlers"
	"github.com/kc1491/go-gql-server/pkg/utils"
)

var HOST, PORT string

func init() {
	HOST = utils.MustGet("GQL_SERVER_HOST")
	PORT = utils.MustGet("GQL_SERVER_PORT")
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + HOST + ":" + PORT)
	log.Fatalln(r.Run(HOST + ":" + PORT))
}
