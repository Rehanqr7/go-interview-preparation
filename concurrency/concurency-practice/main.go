// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// type Order struct {
// 	ID     int
// 	Status string
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	var wg sync.WaitGroup
// 	orderChan := make(chan *Order)
// 	processedOrder := make(chan *Order)

// 	// Add 1 for processOrder goroutine
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		processOrder(orderChan, processedOrder)
// 	}()

// 	// Add 1 for printer goroutine
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for order := range processedOrder {
// 			fmt.Printf("Order processed: ID = %d, Status = %s\n", order.ID, order.Status)
// 		}
// 	}()

// 	// Send orders
// 	for i := 1; i <= 20; i++ {
// 		order := getOrders(i)
// 		orderChan <- order
// 	}
// 	close(orderChan) // Signal no more orders

// 	wg.Wait() // Wait for both goroutines to finish
// }

// func processOrder(orders <-chan *Order, processedOrder chan<- *Order) {
// 	statuses := []string{"Processing", "Delivered", "InTransit"}

// 	for order := range orders {
// 		order.Status = statuses[rand.Intn(len(statuses))]
// 		processedOrder <- order
// 	}

// 	// ✅ Only the writer closes the channel
// 	close(processedOrder)
// }

// func getOrders(id int) *Order {
// 	return &Order{ID: id}
// }

package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println("ch1:", msg)
	case msg := <-ch2:
		fmt.Println("ch2:", msg)
	}
}
