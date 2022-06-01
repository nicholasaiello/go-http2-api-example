package routes

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nicholasaiello/go-http2-api-example/entity"
	"github.com/nicholasaiello/go-http2-api-example/proto"
	v1 "github.com/nicholasaiello/go-http2-api-example/routes/api/v1"
	"google.golang.org/grpc/metadata"
)

var (
	html = template.Must(template.New("").ParseGlob("views/*.html"))
)

type accountHoldingsServer struct {
	proto.UnimplementedAccountHoldingsServer
}

func GetHoldingFromPosition(position *entity.Position) proto.Holding {
	return proto.Holding{
		Id:            position.ID,
		Symbol:        position.Symbol,
		Quantity:      int32(position.Quantity),
		Lots:          []*proto.Lot{},
		PricePaid:     float32(position.PricePaid),
		TotalCost:     float32(position.TotalCost),
		AcquiredDate:  position.AcquiredDate,
		LastTradeDate: position.LastTradeDate,
	}
}

func GetHoldingsFromPortfolio(portfolio *entity.Portfolio) proto.Holdings {
	positions := make([]*proto.Holding, portfolio.TotalPositionsCount)
	for i := 0; i < len(positions); i++ {
		position := GetHoldingFromPosition(&portfolio.Positions[i])
		positions[i] = &position
	}

	return proto.Holdings{
		Id:                  portfolio.ID,
		Positions:           positions,
		TotalPositionsCount: int32(portfolio.TotalPositionsCount),
		ViewName:            portfolio.ViewName,
		LastUpdated:         portfolio.LastUpdated,
	}
}

const DEFAULT_PORTFOLIO_SIZE int = 500

func (s *accountHoldingsServer) GetHoldings(ctx context.Context, account *proto.Account) (*proto.Holdings, error) {
	var totalSize int

	md, ok := metadata.FromIncomingContext(ctx)
	if ok && md["x-portfolio-size"][0] != "" {
		totalSize, _ = strconv.Atoi(md["x-portfolio-size"][0])
	} else {
		totalSize = DEFAULT_PORTFOLIO_SIZE
	}

	portfolio := v1.CreateTestPortfolio(account.Id, totalSize)

	result := GetHoldingsFromPortfolio(&portfolio)

	return &result, nil
}

func (s *accountHoldingsServer) GetHolding(ctx context.Context, holding *proto.Holding) (*proto.Holding, error) {
	position := v1.CreateTestPosition(holding.Id)

	result := GetHoldingFromPosition(&position)

	return &result, nil
}

func (s *accountHoldingsServer) ListHoldings(account *proto.Account, server proto.AccountHoldings_ListHoldingsServer) error {
	var totalSize int

	md, ok := metadata.FromIncomingContext(server.Context())
	if ok && md["x-portfolio-size"][0] != "" {
		totalSize, _ = strconv.Atoi(md["x-portfolio-size"][0])
	} else {
		totalSize = DEFAULT_PORTFOLIO_SIZE
	}

	portfolio := v1.CreateTestPortfolio(account.Id, totalSize)

	positions := make([]proto.Holding, totalSize)
	for i := 0; i < totalSize; i++ {
		positions[i] = GetHoldingFromPosition(&portfolio.Positions[i])

		if err := server.Send(&positions[i]); err != nil {
			return err
		}
	}

	return nil
}

func GetIndex(c *gin.Context) {
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
		if err := pusher.Push("/api/v1/accounts", options); err != nil {
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
		if err := pusher.Push("/api/v1/accounts", options); err != nil {
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
		if err := pusher.Push("/api/v1/accounts", options); err != nil {
			log.Printf("Failed to push: %v\n", err)
		}
	}
	c.HTML(http.StatusOK, "views/unbounded.html", gin.H{
		"status": "success",
		"title":  "Test Unbounded Request",
	})
}

func GetGrpc(c *gin.Context) {
	enablePusher := c.Request.Header.Get("X-Enable-Push") != ""

	if pusher := c.Writer.Pusher(); pusher != nil && enablePusher {
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": c.Request.Header["Accept-Encoding"],
			},
		}
		if err := pusher.Push("/api/v1/accounts", options); err != nil {
			log.Printf("Failed to push: %v\n", err)
		}
	}
	c.HTML(http.StatusOK, "views/grpc.html", gin.H{
		"status": "success",
		"title":  "Test gRPC Request",
	})
}

func GRPCRouter() *accountHoldingsServer {
	return &accountHoldingsServer{}
}

func Router() *gin.Engine {
	router := gin.Default()
	// router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".mp4"})))

	router.SetHTMLTemplate(html)
	router.Static("/assets", "./assets")

	return registerRoutes(router)
}

func registerRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/", GetIndex)
	r.GET("/bounded", GetBounded)
	r.GET("/incremental", GetIncremental)
	r.GET("/unbounded", GetUnbounded)
	r.GET("/grpc", GetGrpc)

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())

	// accounts := apiv1.Group("/accounts")
	apiv1.GET("/accounts", v1.GetAccounts)
	apiv1.GET("/accounts/:id", v1.GetAccount)

	apiv1.GET("/portfolios", v1.GetPortfolios)
	apiv1.GET("/portfolios/:id", v1.GetPortfolio)

	return r
}
