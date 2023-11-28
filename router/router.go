package router

import (
	"github.com/gin-gonic/gin"
)

func Init(){
  //Initialize Router
  r := gin.Default()
  rg := r.Group("/api/v1", )
  //Initialize Routes
  InitRoutes(rg)
  // Run the server 
  _ = r.Run(":8080")
}


