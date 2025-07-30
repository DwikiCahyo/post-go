package main

import (
	"log"

	"github.com/dwikiCahyop/post-go/internal/configs"
	"github.com/dwikiCahyop/post-go/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)



func main(){
	r := gin.Default()

	var (

		cfg *configs.Config

	)


	err := configs.Init(
		// configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config2"),
		// configs.WithConfigType("yaml"),  
	)

	if err != nil {
		log.Fatal("Error read config : " , err)
	}

	cfg = configs.Get()
	log.Println("config" , cfg.Service)

  	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()
	
  
  	r.Run()
}