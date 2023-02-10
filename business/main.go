package main

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

func main() {
	r := gin.Default()

	api := r.Group("/api")

	api.POST("/order", func(c *gin.Context) {
		var order Order
		c.BindJSON(&order)

		res, err := http.Post(
			"http://localhost:8100/api/order",
			"application/json",
			bytes.NewReader([]byte(order.dict)),
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"location": order.Location,
			"time":     order.Time,
			"data":     order.Data,
		})
	})

	api.GET("/record", func(c *gin.Context) {

	})
	r.Run(":8100")
}
