package main

import "fmt"

type Order struct {
	Partner string // Đối tác vận chuyển, ví dụ: "grab", "ghn"
}

// ShipOrder xử lý việc vận chuyển dựa vào đối tác.
func ShipOrder(order Order) {
	switch order.Partner {
	case "grab":
		// ...
		println("Shipping order with Grab.")
	case "ghn":
		// ...
		println("Shipping order with GHN.")
	default:
		println("No valid shipping partner found.")
	}
}

// ShippingStrategy định nghĩa giao diện cho các phương thức vận chuyển.
type ShippingStrategy interface {
	Ship() string
}

// GrabShipping cho phương thức vận chuyển của Grab.
type GrabShipping struct{}

func (GrabShipping) Ship() string {
	return "Order shipped with Grab."
}

// GHNShipping cho phương thức vận chuyển của GHN.
type GHNShipping struct{}

func (GHNShipping) Ship() string {
	return "Order shipped with GHN."
}

// OrderProcessor sử dụng các chiến lược vận chuyển.
type OrderProcessor struct {
	strategy ShippingStrategy
}

func (op *OrderProcessor) SetStrategy(strategy ShippingStrategy) {
	op.strategy = strategy
}

func (op *OrderProcessor) ProcessOrder() string {
	if op.strategy == nil {
		return "No valid shipping strategy set."
	}
	return op.strategy.Ship()
}

func main() {
	order := Order{Partner: "grab"}
	processor := &OrderProcessor{}
	// ...
	switch order.Partner {
	case "grab":
		processor.SetStrategy(GrabShipping{})
	case "ghn":
		processor.SetStrategy(GHNShipping{})
	default:
		fmt.Println("No valid shipping partner found.")
		return
	}
	fmt.Println(processor.ProcessOrder())
}

// func main() {
// 	order1 := Order{Partner: "Grab"}
// 	ShipOrder(order1)
// }
