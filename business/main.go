package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"os"
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

func main() {
	r := gin.Default()

	api := r.Group("/api")

	api.POST("/order", func(c *gin.Context) {

		var order Order
		c.BindJSON(&order)

		values := map[string]interface{}{
			"location":  order.Location,
			"timestamp": order.Time,
			"data":      order.Data,
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

		//todo 資料

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	api.GET("/record", func(c *gin.Context) {
		SQL_HOST := os.Getenv("SQL_HOST")
		SQL_PORT := os.Getenv("SQL_PORT")
		log.Println(SQL_HOST)
		log.Println(SQL_PORT)
		db, err := sql.Open("postgres", "postgresql://user:123456@"+SQL_HOST+":"+SQL_PORT+"/tsmc_pg")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		m := "Steel"
		rows, err := db.Query("SELECT * FROM data WHERE material = $1", m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		materials := []string{}

		for rows.Next() {
			var material string
			if err := rows.Scan(&material); err != nil {
				materials = append(materials, material)
			}
		}
		c.JSON(200, gin.H{
			"message": "success",
			"data":    materials,
		})
	})

	api.GET("/report", func(c *gin.Context) {

	})

	r.Run(":8100")
}
