package main

import (
	"fmt"
	"time"
	"encoding/base64"
	"github.com/gin-gonic/gin"
)

type Inventory struct {
	Material uint64      `json:"material"`
	Signature string  `json:"signature"`
}

type Order struct {
	Location string    `json:"location"`
	Time     time.Time `json:"timestamp"`
	Data     Data      `json:"data"`
}


type Data struct {
	A uint64 `json:"a"`
	B uint64 `json:"b"`
	C uint64 `json:"c"`
	D uint64 `json:"d"`
}

func _signature(total uint64) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d",total)))
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	
	api.POST("/order", func(c *gin.Context) {
		var order Order
		c.BindJSON(&order)
		fmt.Println(order)
		var inventory Inventory
		total := order.Data.A + order.Data.B + order.Data.C + order.Data.D
		inventory.Material = order.Data.A * 3 + order.Data.B * 2 + order.Data.C * 4 + order.Data.D * 10
		inventory.Signature = _signature(total)
		fmt.Println(inventory)
		fmt.Println(total)
		c.JSON(200, gin.H{
			"material": inventory.Material,
			"signature": inventory.Signature,
		})
	})
	r.Run(":8200") // listen and serve on
}