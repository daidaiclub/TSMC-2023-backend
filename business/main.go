package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type Order struct {
	Location string    `json:"location"`
	Time     time.Time `json:"timestamp"`
	Data     Data      `json:"data"`
}

type Inventory struct {
	Material  uint64 `json:"material"`
	Signature string `json:"signature"`
}

type Record struct {
	Location  string    `json:"location"`
	Time      time.Time `json:"timestamp"`
	Material  uint64    `json:"material"`
	Signature string    `json:"signature"`
	Data      Data      `json:"data"`
}

type Data struct {
	A uint64 `json:"a"`
	B uint64 `json:"b"`
	C uint64 `json:"c"`
	D uint64 `json:"d"`
}

// func OrderToByteArray(Order order) []byte {

// }


var flag1, flag2 bool
func main() {
	r := gin.Default()
	flag1 = true
	flag2 = true
	api := r.Group("/api")

	api.POST("/order", func(c *gin.Context) {
		inventoryEnd := true
		var order Order
		c.BindJSON(&order)
		values := map[string]interface{}{
			"location":  order.Location,
			"timestamp": order.Time,
			"data":      order.Data,
		}

		jsonValue, _ := json.Marshal(values)
		var res *http.Response
		var err error
		for inventoryEnd {
			res, err = http.Post(
				"http://localhost:8200/api/order",
				"application/json",
				bytes.NewBuffer(jsonValue),
			)

			if err != nil {
				time.Sleep(20 * time.Second)
				continue
			}
			inventoryEnd = false
		}

		var inventory Inventory
		sitemap, err := io.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sitemap_s := []byte(string(sitemap))
		json.Unmarshal(sitemap_s, &inventory)


		var Record Record
		Record.Location = order.Location
		Record.Time = order.Time
		Record.Material = inventory.Material
		Record.Signature = inventory.Signature
		Record.Data = order.Data
	
		values = map[string]interface{}{
			"location":  Record.Location,
			"timestamp": Record.Time,
			"material":  Record.Material,
			"signature": Record.Signature,
			"data":      Record.Data,
		}
		
		jsonValue, _ = json.Marshal(values)
		storageEnd := true
		for storageEnd {
			res, err = http.Post(
				"http://localhost:8300/api/records",
				"application/json",
				bytes.NewBuffer(jsonValue),
			)

			if err != nil {
				time.Sleep(20 * time.Second)
				continue
			}
			storageEnd = false
		}
		
		//todo 資料

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	api.GET("/record", func(c *gin.Context) {
		for flag2 {
			flag1 = false
			res, err := http.Get("http://localhost:8100/api/check")
			if err != nil {
				time.Sleep(50 * time.Millisecond)
				continue
			}
			dataMap, err := io.ReadAll(res.Body)
			if err != nil {
				time.Sleep(50 * time.Millisecond)
				continue
			}
			dataMap_s := []byte(string(dataMap))
			var dataMap_m map[string]interface{}
			json.Unmarshal(dataMap_s, &dataMap_m)
			flag2 = dataMap_m["flag"].(bool)
			time.Sleep(50 * time.Millisecond)
		}
	})

	api.GET("/report", func(c *gin.Context) {
		for flag2 {
			flag1 = false
			res, err := http.Get("http://localhost:8100/api/check")
			if err != nil {
				continue
			}
			dataMap, err := io.ReadAll(res.Body)
			if err != nil {
				continue
			}
			dataMap_s := []byte(string(dataMap))
			var dataMap_m map[string]interface{}
			json.Unmarshal(dataMap_s, &dataMap_m)
			flag2 = dataMap_m["flag"].(bool)
			time.Sleep(50 * time.Millisecond)
		}
	})

	api.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"flag": flag1,
		})
	})

	r.Run(":8100")
}
