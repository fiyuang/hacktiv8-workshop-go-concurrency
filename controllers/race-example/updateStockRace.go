package race_example

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

type Product struct {
	name  string
	stock int
}

func NewProduct(name string, initialStock int) *Product {
	return &Product{name: name, stock: initialStock}
}

func (s *Product) AddStock(amount int) {
	s.stock += amount
	fmt.Printf("Added %d Product to %s. Current stock: %d\n", amount, s.name, s.stock)
}

func (s *Product) SellProduct(amount int) {
	if s.stock < amount {
		fmt.Printf("%s doesn't have enough stock. You can't sell with this quantity %d. Because the current stock: %d\n", s.name, amount, s.stock)
		return
	}
	s.stock -= amount
	fmt.Printf("Sold %d Product from %s. Current stock: %d\n", amount, s.name, s.stock)
}

func UpdateStock(ctx *gin.Context) {
	product := NewProduct("Tas Ransel", 1000)
	totalUser := 100

	for i := 0; i < totalUser; i++ {
		go func() {
			amount := rand.Intn(10)
			product.SellProduct(amount)
			return
		}()
	}

	ctx.JSON(200, gin.H{
		"product_name": product.name,
		"last_stocks":  product.stock,
	})
}
