package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nicholasaiello/go-middleware-example/entity"
)

func CreateTestAccount(description string, mode string, accountStatus string, accountType string) entity.Account {
	id := CreateRandID(4)

	account := entity.Account{
		ID:          id,
		EncodedID:   CreateHash(id),
		Name:        description + " " + id[4:] + "-" + id[:4],
		ShortName:   description + " -" + id[:4],
		Description: description,
		Mode:        mode,
		Status:      accountStatus,
		Type:        accountType,
		Created:     "11/28/2017",
	}

	return account
}

func CreateTestBrokerageAccount() entity.Account {
	return CreateTestAccount("Individual Brokerage", "CASH", "ACTIVE", "INDIVIDUAL")
}

func CreateTestAccounts(size int) []entity.Account {
	var accounts = make([]entity.Account, size)
	for i := 0; i < size; i++ {
		accounts[i] = CreateTestBrokerageAccount()
	}
	return accounts
}

var (
	testAccounts = CreateTestAccounts(5)
)

func GetAccounts(c *gin.Context) {
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
