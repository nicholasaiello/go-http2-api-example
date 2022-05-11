package routes

import (
	"html/template"
	"log"
	"net/http"
	v1 "nicholasaiello/go-middleware-example/routes/api/v1"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()

	html = template.Must(template.New("").ParseGlob("views/*.html"))
)

func GetIndex(c *gin.Context) {
	pushRequestHeader := c.Request.Header.Get("X-Enable-Push")
	enablePusher := pushRequestHeader == "chunked" || pushRequestHeader == "push"

	if pusher := c.Writer.Pusher(); pusher != nil && enablePusher {
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": c.Request.Header["Accept-Encoding"],
			},
		}
		if err := pusher.Push("/api/v1/accounts/", options); err != nil {
			log.Printf("Failed to push: %v\n", err)
		}
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{"status": "success"})
}

func Run(address string) {
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".mp4"})))

	router.SetHTMLTemplate(html)
	router.Static("/assets", "./assets")

	createRoutes()

	router.RunTLS(address, "./server.crt", "./server.key")
}

func createRoutes() {
	router.GET("/", GetIndex)

	apiv1 := router.Group("/api/v1")
	// apiv1.Use(jwt.JWT())

	// accounts := apiv1.Group("/accounts")
	apiv1.GET("/accounts", v1.GetAccounts)
	apiv1.GET("/accounts/:id", v1.GetAccount)

	apiv1.GET("/portfolios", v1.GetPortfolios)
	apiv1.GET("/portfolios/:id", v1.GetPortfolio)
}
