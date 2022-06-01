package v1

import "github.com/nicholasaiello/go-http2-api-example/entity"

func CreateTestPosition(id string) entity.Position {
	if id == "" {
		id = CreateRandID(4)
	}

	symbol := GetRandomSymbol()
	quantity := CreateRandInt(10000) + 1
	pricePaid := CreateRandFloat(500)

	position := entity.Position{
		ID:            id,
		Symbol:        symbol,
		Details:       entity.ProductDetails{},
		Quantity:      int64(quantity),
		Lots:          []entity.PositionLot{},
		PricePaid:     pricePaid,
		TotalCost:     pricePaid * float64(quantity),
		Fees:          0,
		AcquiredDate:  "10/03/2017",
		LastTradeDate: "04:00 PM ET 05/06/22",
	}

	return position
}

func CreateTestPositions(size int) []entity.Position {
	positions := make([]entity.Position, size)
	for i := 0; i < size; i++ {
		positions[i] = CreateTestPosition("")
	}
	return positions
}

func CreateTestPortfolio(id string, totalPositions int) entity.Portfolio {
	if id == "" {
		id = CreateRandID(4)
	}

	positions := CreateTestPositions(totalPositions)

	portfolio := entity.Portfolio{
		ID:                  id,
		Positions:           positions,
		TotalPositionsCount: len(positions),
		ViewName:            "All Positions",
		LastUpdated:         "04:00 PM ET 05/06/22",
	}

	return portfolio
}
