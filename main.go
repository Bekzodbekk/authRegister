package main

import (
	"context"
	"fmt"
	"log"
	"service/config"
	"service/models"
	"service/mongosh"
	"service/storage"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerSt struct {
	User *storage.UserRepo
}

func (h *HandlerSt) SignUp(c *gin.Context) {

	req := models.User{}
	err := c.BindJSON(&req)
	fmt.Println(req)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = h.User.SignUp(&req)
	if err != nil {
		c.JSON(400, gin.H{"message1": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "succesfully"})
}

func NewGin(c *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// Setup CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins (change this to specific origins for production)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	handler := HandlerSt{
		User: storage.NewUserRepo(c),
	}

	r.POST("/sign-up", handler.SignUp)

	return r
}

func main() {

	cfg, err := config.Load("./")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	m := &mongosh.MongoDB{}
	go func() {
		defer wg.Done()
		m, err = mongosh.GetCollection(*cfg)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()

	r := NewGin(&m.Collection)
	defer m.Client.Disconnect(context.Background())

	r.Run()
}
