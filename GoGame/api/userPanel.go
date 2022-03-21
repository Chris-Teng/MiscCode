package api

import (
	"game/config"
	"log"

	"github.com/gin-gonic/gin"
)

func QueryStatistics(c *gin.Context) {
	conn := config.Dbconnect()

	rows, err := conn.Query("SELECT * FROM statistics")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//一个var声明多个变量的操作
	var (
		id               int16
		name_of_power    string
		num_of_cities    int32
		total_population int32
		total_army       int32
	)

	for rows.Next() {
		if err := rows.Scan(&id, &name_of_power, &num_of_cities, &total_population, &total_army); err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{
			"id":               id,
			"name_of_power":    name_of_power,
			"num_of_cities":    num_of_cities,
			"total_population": total_population,
			"total_army":       total_army,
		})
	}
}

// func UpdateAllStatus
