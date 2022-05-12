package v1

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicholasaiello/go-http2-api-example/entity"
)

const DEFAULT_LIMIT_QUERY = "100"
const DEFAULT_PORTFOLIO_SIZE = 5000

var (
	testPortfolios = []entity.Portfolio{
		CreateTestPortfolio("", 1000),
	}
)

func fetchPortfolio(accountId string, portfolioSize int) entity.Portfolio {
	time.Sleep(100 * time.Millisecond)

	return CreateTestPortfolio(accountId, portfolioSize)
}

func GetPortfolios(c *gin.Context) {
	c.JSON(http.StatusOK, testPortfolios)
}

func GetPortfolio(c *gin.Context) {
	accountId := c.Param("id")

	chunkQuery := c.Query("chunk")
	limitQuery := c.Query("limit")
	startQuery := c.Query("start")

	// FIXME: hack to toggle Push mode
	pushRequestHeader := c.Request.Header.Get("X-Enable-Push")
	enablePusher := pushRequestHeader == "chunked" || pushRequestHeader == "push"

	// FIXME: hack to modify portfolio size
	portfolioSizeHeader := c.Request.Header.Get("X-Portfolio-Size")
	if portfolioSizeHeader == "" {
		portfolioSizeHeader = "1000"
	}

	portfolioSize, _ := strconv.Atoi(portfolioSizeHeader)
	testPortfolio := fetchPortfolio(accountId, portfolioSize)

	// Unbounded
	if limitQuery == "" && startQuery == "" {
		c.JSON(http.StatusOK, testPortfolio)
		return
	}

	testBoundedPortfolio := testPortfolio
	totalPositions := len(testPortfolio.Positions)

	// Apply defaults for bounds
	if limitQuery == "" {
		limitQuery = DEFAULT_LIMIT_QUERY
	}
	if startQuery == "" {
		startQuery = "0"
	}

	limit, _ := strconv.Atoi(limitQuery)
	start, _ := strconv.Atoi(startQuery)

	// Use the smaller limit
	if (totalPositions - start) < limit {
		limit = totalPositions - start
	}

	// Simple defensive checks
	if (limit <= 0 || start < 0) || start >= totalPositions {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid boundaries"})
		return
	}

	shouldCreatePusherRequest := enablePusher && start == 0

	// Optimized push request(s) if `Writer.Pusher` is supported, and starting from first page
	if pusher := c.Writer.Pusher(); pusher != nil && shouldCreatePusherRequest {
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": c.Request.Header["Accept-Encoding"],
			},
		}

		if pushRequestHeader == "chunked" {
			if chunkQuery == "" {
				chunkQuery = limitQuery
			}

			chunk, _ := strconv.Atoi(chunkQuery)
			chunks := float64(totalPositions-limit) / float64(chunk)

			log.Printf("Creating preload request for %d chunks", int(math.Round(chunks)))

			CreateChunkedPortfolioPushRequests(pusher, "/api/v1/portfolios/"+accountId, start+limit, chunk, int(math.Round(chunks)), options)
		} else {
			CreatePortfolioPushRequest(pusher, "/api/v1/portfolios/"+accountId, start+limit, totalPositions-limit, options)
		}
	}

	log.Printf("Request for portfolio, start = %d, limit = %d", start, limit)

	testBoundedPortfolio.Positions = testBoundedPortfolio.Positions[start : start+limit]

	c.JSON(http.StatusOK, testBoundedPortfolio)
}

func CreatePortfolioPushRequest(pusher http.Pusher, url string, start int, limit int, options *http.PushOptions) {
	newUrl := url + "?start=" + strconv.Itoa(start) + "&limit=" + strconv.Itoa(limit)
	if err := pusher.Push(newUrl, options); err != nil {
		log.Printf("Failed to push: %v\n", err)
	}
}

func CreateChunkedPortfolioPushRequests(pusher http.Pusher, url string, start int, chunk int, chunks int, options *http.PushOptions) {
	chunkQuery := strconv.Itoa(chunk)

	for i := 0; i < chunks; i++ {
		newStart := strconv.Itoa(start + (chunk * i))
		newUrl := url + "?start=" + newStart + "&limit=" + chunkQuery

		if err := pusher.Push(newUrl, options); err != nil {
			log.Printf("Failed to push: %v\n", err)
			break
		}
	}
}
