package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	v1 "github.com/nicholasaiello/go-http2-api-example/routes/api/v1"
)

var (
	html = template.Must(template.New("").ParseGlob("views/*.html"))
)

func GetIndex(c *gin.Context) {
	enablePusher := c.Request.Header.Get("X-Enable-Push") != ""

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
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"status": "success",
		"title":  "Test",
	})
}

func GetBounded(c *gin.Context) {
	enablePusher := c.Request.Header.Get("X-Enable-Push") != ""

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
	c.HTML(http.StatusOK, "views/bounded.html", gin.H{
		"status": "success",
		"title":  "Test Bounded Request",
	})
}

func GetIncremental(c *gin.Context) {
	enablePusher := c.Request.Header.Get("X-Enable-Push") != ""

	if pusher := c.Writer.Pusher(); pusher != nil && enablePusher {
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": c.Request.Header["Accept-Encoding"],
				"X-Enable-Push":   c.Request.Header["X-Enable-Push"],
			},
		}
		if err := pusher.Push("/api/v1/accounts/", options); err != nil {
			log.Printf("Failed to push: %v\n", err)
		}
	}
	c.HTML(http.StatusOK, "views/incremental.html", gin.H{
		"status": "success",
		"title":  "Test Incremental Request",
	})
}

func GetUnbounded(c *gin.Context) {
	enablePusher := c.Request.Header.Get("X-Enable-Push") != ""

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
	c.HTML(http.StatusOK, "views/unbounded.html", gin.H{
		"status": "success",
		"title":  "Test Unbounded Request",
	})
}

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".mp4"})))

	router.SetHTMLTemplate(html)
	router.Static("/assets", "./assets")

	createRoutes(router)

	return router
}

func createRoutes(r *gin.Engine) {
	r.GET("/", GetIndex)
	r.GET("/bounded", GetBounded)
	r.GET("/incremental", GetIncremental)
	r.GET("/unbounded", GetUnbounded)

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())

	// accounts := apiv1.Group("/accounts")
	apiv1.GET("/accounts", v1.GetAccounts)
	apiv1.GET("/accounts/:id", v1.GetAccount)

	apiv1.GET("/portfolios", v1.GetPortfolios)
	apiv1.GET("/portfolios/:id", v1.GetPortfolio)
}
