package main

import (
	"log"
	"time"

	"5ub5i5t/bughuntingtoolbox/controller"
	"5ub5i5t/bughuntingtoolbox/database"
	"5ub5i5t/bughuntingtoolbox/mitmproxy"
	"5ub5i5t/bughuntingtoolbox/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Domain{})
	//database.Database.AutoMigrate(&proxy.Flow{})
	database.Database.AutoMigrate(&model.CustomFlow{})
}

func serveApplication() {
	router := gin.Default()
	hour := time.Hour

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin,access-control-allow-headers"},
		//AllowHeaders:     []string{"Origin"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * hour,
	}))

	domainsRoutes := router.Group("/api/domains")
	domainsRoutes.GET("/", controller.GetDomains)

	domainRoutes := router.Group("/api/domain")
	domainRoutes.GET("/:id", controller.GetDomainById)
	domainRoutes.POST("/add", controller.AddDomain)
	domainRoutes.PUT("/update/:id", controller.UpdateDomainById)
	domainRoutes.DELETE("/delete/:id", controller.DeleteDomainById)

	proxyRoutes := router.Group("/api/proxy")
	proxyRoutes.GET("/start", mitmproxy.StartProxy)

	go func() {
		mitmproxy.StartProxyBasic()
	}()

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
