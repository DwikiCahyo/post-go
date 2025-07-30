package main

import (
	"log"

	"github.com/dwikiCahyop/post-go/internal/configs"
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()

	var (

		cfg *configs.Config
		databaseName string

	)


	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),  
	)

	if err != nil {
		log.Fatal("Error read config : " , err)
	}

	cfg = configs.Get()


	databaseName = cfg.Database.DataSourceName

	log.Println("get config : " , cfg)
	log.Println("database" , databaseName)


	// user := new(User)

	// err = db.QueryRow("SELECT id, email FROM users WHERE id = ? ", 1).Scan(&user.id, &user.email)
	// if err != nil {
	// 	log.Fatal("error query database : ", err)
	// }

	// log.Println("user: ", user)

  	// membershipHandler := memberships.NewHandler(r)
	// membershipHandler.RegisterRoute()
	
  
  	r.Run()
}