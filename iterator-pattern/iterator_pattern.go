package main

import "fmt"

type Order struct {
	ID    string
	Items []Item
}

type Item struct {
	Name     string
	Price    int // Giá của món
	Discount int // Giảm giá được áp dụng trên món
}

func createMapFromListItems(items []Item) map[string]Item {
	itemsMap := make(map[string]Item)
	for _, item := range items {
		itemsMap[item.Name] = item
	}
	return itemsMap
}

func calculateTotalPriceArray(items []Item) int {
	total := 0
	for _, item := range items {
		total += item.Price
	}
	return total
}

func calculateTotalDiscountArray(items []Item) int {
	totalDiscount := 0
	for _, item := range items {
		totalDiscount += item.Discount
	}
	return totalDiscount
}

func calculateTotalPriceMap(items map[string]Item) int {
	total := 0
	for _, item := range items {
		total += item.Price
	}
	return total
}

func calculateTotalDiscountMap(items map[string]Item) int {
	totalDiscount := 0
	for _, item := range items {
		totalDiscount += item.Discount
	}
	return totalDiscount
}

// Implement Iterator Pattern
// Iterator interface định nghĩa phương thức cần thiết để duyệt qua các phần tử.
type Iterator interface {
	HasNext() bool
	Next() *Item
	Reset() // just testing
}

// Concrete Iterator cho slice
type sliceIterator struct {
	items []Item
	index int
}

func (it *sliceIterator) HasNext() bool {
	return it.index < len(it.items)
}

func (it *sliceIterator) Next() *Item {
	if it.HasNext() {
		item := &it.items[it.index]
		it.index++
		return item
	}
	it.index = 0
	return nil
}

func (it *sliceIterator) Reset() {
	it.index = 0
}

// Concrete Iterator cho map[string]Item
type mapIterator struct {
	items map[string]Item
	keys  []string
	index int
}

func (it *mapIterator) HasNext() bool {
	return it.index < len(it.keys)
}

func (it *mapIterator) Next() *Item {
	if it.HasNext() {
		key := it.keys[it.index]
		item := it.items[key]
		it.index++
		return &item
	}
	it.index = 0
	return nil
}

func (it *mapIterator) Reset() {
	it.index = 0
}

// Aggregate interface định nghĩa phương thức tạo iterator.
type Aggregate interface {
	CreateIterator() Iterator
}

// Concrete Aggregate cho slice
type ItemSlice struct {
	items []Item
}

func (a *ItemSlice) CreateIterator() Iterator {
	return &sliceIterator{items: a.items}
}

// Concrete Aggregate cho map
type ItemMap struct {
	items map[string]Item
}

func (m *ItemMap) CreateIterator() Iterator {
	keys := make([]string, 0, len(m.items))
	for key := range m.items {
		keys = append(keys, key)
	}
	return &mapIterator{items: m.items, keys: keys}
}

func calculateTotalPrice(iterator Iterator) int {
	total := 0
	for iterator.HasNext() {
		total += iterator.Next().Price
	}
	return total
}

func calculateTotalDiscount(iterator Iterator) int {
	total := 0
	for iterator.HasNext() {
		total += iterator.Next().Discount
	}
	return total
}

func main() {
	order := Order{
		ID: "24ABCXYZ",
		Items: []Item{
			{"Kem", 50000, 5000},
			{"Bánh", 10000, 0},
			{"Cà phê", 70000, 10000},
		},
	}
	// Slice
	itemSlices := ItemSlice{items: order.Items}
	iteratorSlice := itemSlices.CreateIterator()
	fmt.Printf("Total price: %d\n", calculateTotalPrice(iteratorSlice))
	iteratorSlice.Reset()
	fmt.Printf("Total discount: %d\n", calculateTotalDiscount(iteratorSlice))

	// Map
	itemsMap := createMapFromListItems(order.Items)
	itemsMapAggregate := ItemMap{items: itemsMap}
	iteratorMap := itemsMapAggregate.CreateIterator()
	fmt.Printf("Total price: %d\n", calculateTotalPrice(iteratorMap))
	iteratorMap.Reset()
	fmt.Printf("Total discount: %d\n", calculateTotalDiscount(iteratorMap))
}

func CreateSliceIterator(items []Item) Iterator {
	return &sliceIterator{items: items}
}
