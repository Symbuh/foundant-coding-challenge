package main

import (
	"net/http"
	// I have to replace this with the filepath to my created api package
	//"github.com/gowebexamples/http-server/api"
	// I guess this should be pointing to my github instead of what it's pointing to
	"github.com/Symbuh/foundand-coding-challenge/api"
	// "github.com/gin-gonic/contrib/static"
	// "github.com/gin-gonic/gin"
)

func main() {

	server := api.NewServer()

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", server)

	// router := gin.Default()

	// router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// api := router.Group("/api")
	// {
	// 	api.GET("/", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})
	// }

	// router.Run(":5000")
}
