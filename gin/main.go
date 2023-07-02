package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		dataStr, err := os.ReadFile("../data/pokemon.json", )

		if err != nil {
			fmt.Println("Error occurred", err)
		}

		var array []map[string]string
		json.Unmarshal(dataStr, &array)

		for _, pokemon := range array {
			if pokemon["Legendary"] == "True" {
				pokemon["Legendary"] = "true"
			} else {
				pokemon["Legendary"] = "false"
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"data": array,
		})
	})
	r.Run()
}