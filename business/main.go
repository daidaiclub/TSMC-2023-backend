package main

import (
	"bytes"
	"net/http"
	"time"
	"io"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Order struct {
	Location string    `json:"location"`
	Time     time.Time `json:"timestamp"`
	Data     Data      `json:"data"`
}

type Inventory struct {
	Material uint64      `json:"material"`
	Signature string  `json:"signature"`
}

type Record struct {
	Location string    `json:"location"`
	Time     time.Time `json:"timestamp"`
	Material uint64       `json:"material"`
	Signature string   `json:"signature"`
	Data     Data      `json:"data"`
}

type Data struct {
	A uint64 `json:"a"`
	B uint64 `json:"b"`
	C uint64 `json:"c"`
	D uint64 `json:"d"`
}


// func OrderToByteArray(Order order) []byte {
   
// }

func main() {
	r := gin.Default()

	api := r.Group("/api")

	
	api.POST("/order", func(c *gin.Context) {
		
		var order Order
		c.BindJSON(&order)


		values := map[string]interface{}{
			"location": order.Location,
			"timestamp": order.Time,
			"data": order.Data,
		}

		jsonValue, _ := json.Marshal(values)

		res, err := http.Post(
			"http://localhost:8200/api/order",
			"application/json",
			bytes.NewBuffer(jsonValue),
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var inventory Inventory
		sitemap, err := io.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		sitemap_s := []byte(string(sitemap))
		json.Unmarshal(sitemap_s, &inventory)

		var Record Record
		Record.Location = order.Location
		Record.Time = order.Time
		Record.Material = inventory.Material
		Record.Signature = inventory.Signature
		Record.Data = order.Data

		
		


		c.JSON(200, gin.H{
			"location": order.Location,
			"time":     order.Time,
			"Record":     Record,
		})
	})

	api.GET("/record", func(c *gin.Context) {

	})

	api.GET("/report", func(c *gin.Context) {
	
	})

	r.Run(":8100")
}
