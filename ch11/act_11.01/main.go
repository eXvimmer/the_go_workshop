package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

type item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

type order struct {
	TotalPrice  int    `json:"total"`
	IsPaid      bool   `json:"paid"`
	Fragile     bool   `json:",omitempty"`
	OrderDetail []item `json:"orderdetail"`
}

type customer struct {
	UserName      string  `json:"username"`
	Password      string  `json:"-"`
	Token         string  `json:"-"`
	ShipTo        address `json:"shipto"`
	PurchaseOrder order   `json:"order"`
}

func (c *customer) total() {
	total := 0
	for _, v := range c.PurchaseOrder.OrderDetail {
		total += v.Quantity * v.Price
	}
	c.PurchaseOrder.TotalPrice = total
}

func (c *customer) placeOrder(name, description string, quantity, price int) {
	i := item{
		Name:        name,
		Description: description,
		Quantity:    quantity,
		Price:       price,
	}

	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, i)
	c.total()
}

func main() {
	jsonData := []byte(`
	{
		"username" :"exvimmer",
		"shipto":  
			{
		    		"street": "main street",
	    			"city": "New York",
	    			"state": "NY",
	    			"zipcode": 12345

			},
		"order":
			{
				"paid": false,
				"orderdetail" : 
					[{
						"itemname":"Pizza",
						"desc": "food",
						"qty": 2,
						"price": 15
					}]
			}
	}
	`)

	if !json.Valid(jsonData) {
		fmt.Printf("invalid json: %s\n", jsonData)
		os.Exit(1)
	}

	var c customer

	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c.placeOrder("Harry Potter and the philosopher's stone", "A book", 2, 20)
	c.placeOrder("Desert Eagle", "A gun", 1, 999)
	c.placeOrder("Mechanical keyboard", "keyboard", 1, 120)

	c.PurchaseOrder.Fragile = true
	c.PurchaseOrder.IsPaid = true

	theOrder, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", theOrder)
}
