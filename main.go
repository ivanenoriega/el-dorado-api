package main

import (
	"log"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/ivanenoriega/el-dorado-api/db"
	m "github.com/ivanenoriega/el-dorado-api/mapper/email"
	mo "github.com/ivanenoriega/el-dorado-api/model/email"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	router.POST("/email", func(c *gin.Context) {
		var data mo.ReqEmail
		con := db.GetConn()
		if err := c.Bind(&data); err != nil {
			c.JSON(400, gin.H{"error": "invalid email"})
			return
		}
		if data.Email == "" {
			c.JSON(400, gin.H{"error": "empty email"})
			return
		}
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if re.MatchString(data.Email) == false {
			c.JSON(400, gin.H{"error": "malformed email"})
			return
		}
		err := m.Insert(data.Email, con)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(200, gin.H{
			"message": "Success",
		})
		defer con.Close()
	})

	router.Run(":" + port)
}
