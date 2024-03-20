package main

import (
	"cinema/component/appctx"
	"cinema/middleware"
	gincinema "cinema/module/cinema/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Read the content of the db_env file
	content, err := ioutil.ReadFile("local_db_env")
	if err != nil {
		log.Fatalln(err)
	}
	// Replace newline characters with spaces
	dsn := strings.ReplaceAll(string(content), "\n", " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.Debug()

	appCtx := appctx.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	cinemas := v1.Group("/cinemas")
	//GET /v1/cinemas
	cinemas.GET("", gincinema.ListCinema(appCtx))

	//POST /v1/cinemas
	cinemas.POST("", gincinema.CreateCinema(appCtx))

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
