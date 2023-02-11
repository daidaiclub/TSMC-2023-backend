package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	Location  string `json:"location"`
	Time      string `json:"timestamp"`
	Material  uint64 `json:"material"`
	Signature string `json:"signature"`
	Data      Data   `json:"data"`
}

type Data struct {
	A uint64 `json:"a"`
	B uint64 `json:"b"`
	C uint64 `json:"c"`
	D uint64 `json:"d"`
}

type Report struct {
	Location string `json:"location"`
	Date     string `json:"date"`
	Count    uint64 `json:"count"`
	Material uint64 `json:"material"`
	A        uint64 `json:"a"`
	B        uint64 `json:"b"`
	C        uint64 `json:"c"`
	D        uint64 `json:"d"`
}

// func OrderToByteArray(Order order) []byte {

// }

func get_db() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://user:123456@postgres:5432/tsmc-storage?sslmode=disable")
	return db, err
}

func main() {
	db, err := get_db()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create data table if not exist
	_, table_check := db.Query("SELECT * FROM data")
	if table_check != nil {
		_, err = db.Exec(`
			CREATE TABLE data (
				location text,
				timestamp text,
				date text,s
				material integer,
				signature text,
				A integer,
				B integer,
				C integer,
				D integer);`,
		)
		if err != nil {
			panic(err)
		}
	}

	r := gin.Default()

	api := r.Group("/api")

	api.POST("/records", func(c *gin.Context) {
		var record Record
		c.BindJSON(&record)

		_, err = db.Exec("INSERT INTO data (location, timestamp, date, material, signature, a, b, c, d) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			record.Location, record.Time, record.Time[:10], record.Material, record.Signature, record.Data.A, record.Data.B, record.Data.C, record.Data.D)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	api.GET("/records", func(c *gin.Context) {
		location := c.Query("location")
		date := c.Query("date")

		rows, err := db.Query("SELECT * FROM data WHERE location = $1 AND date = $2", location, date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var records []Record
		for rows.Next() {
			var record Record
			err = rows.Scan(&record.Location, &record.Time, &record.Material, &record.Signature, &record.Data.A, &record.Data.B, &record.Data.C, &record.Data.D)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			records = append(records, record)
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    records,
		})
	})

	api.GET("/report", func(c *gin.Context) {
		location := c.Query("location")
		date := c.Query("date")
		rows, err := db.Query(
			`SELECT location, date, COUNT(*) AS count, SUM(material) AS material, SUM(a) AS a, SUM(b) AS b, SUM(c) AS c, SUM(d) AS d
				FROM data 
				WHERE location = $1 AND date = $2
				GROUP BY location, date`, location, date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var report Report
		for rows.Next() {
			err = rows.Scan(&report.Location, &report.Date, &report.Count, &report.Material, &report.A, &report.B, &report.C, &report.D)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(200, gin.H{
			"message":  "success",
			"location": report.Location,
			"date":     report.Date,
			"count":    report.Count,
			"material": report.Material,
			"a":        report.A,
			"b":        report.B,
			"c":        report.C,
			"d":        report.D,
		})
	})

	api.POST("/clean", func(c *gin.Context) {
		_, err = db.Exec("DELETE FROM data")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run(":8300")
}
