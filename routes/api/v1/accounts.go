package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	testAccounts = CreateTestAccounts(5)
)

func GetAccounts(c *gin.Context) {
	includePortfolio := c.Query("includePortfolio") != ""

	if pusher := c.Writer.Pusher(); pusher != nil && includePortfolio {
		selectedAccount := testAccounts[0]
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding":  c.Request.Header["Accept-Encoding"],
				"X-Portfolio-Size": c.Request.Header["X-Portfolio-Size"],
			},
		}

		portfolioUrl := "/api/v1/portfolios/" + selectedAccount.ID
		if err := pusher.Push(portfolioUrl, options); err != nil {
			log.Printf("Failed to push: %v\n", err)
		}
	}

	c.JSON(http.StatusOK, testAccounts)
}

func GetAccount(c *gin.Context) {
	id := c.Param("id")

	for _, a := range testAccounts {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "account not found"})
}
