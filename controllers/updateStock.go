package controllers

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/gin-gonic/gin"
)

type Product struct {
	name  string
	stock int
	mutex sync.Mutex
}

func NewProduct(name string, initialStock int) *Product {
	return &Product{name: name, stock: initialStock}
}

// AddStock safely adds to the stock count
func (s *Product) AddStock(amount int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.stock += amount
	fmt.Printf("Added %d Product to %s. Current stock: %d\n", amount, s.name, s.stock)
}

// SellProduct safely reduces the stock count
func (s *Product) SellProduct(amount int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.stock < amount {
		fmt.Printf("%s doesn't have enough stock. Current stock: %d\n", s.name, s.stock)
		return
	}
	s.stock -= amount
	fmt.Printf("Sold %d Product from %s. Current stock: %d\n", amount, s.name, s.stock)
}

func UpdateStock(ctx *gin.Context) {
	product := NewProduct("Tas Ransel", 1000)
	totalUser := 100

	// Simulate concurrent transactions
	var wg sync.WaitGroup
	wg.Add(totalUser)

	for i := 0; i < totalUser; i++ {

		// if you uncomment this block, update the wg.Add to wg.Add(totalUser*2)
		/* go func() {
			amount := rand.Intn(10)
			product.AddStock(amount)
			wg.Done()
		}() */

		go func() {
			amount := rand.Intn(10)
			product.SellProduct(amount)
			wg.Done()
		}()

	}

	wg.Wait()

	ctx.JSON(200, gin.H{
		"product_name": product.name,
		"last_stocks":  product.stock,
	})
}
