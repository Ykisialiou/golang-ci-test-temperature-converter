package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

type temperature struct {
     TemperatureC string `json:"temperatureC"`
     TemperatureF string `json:"temperatureF"`
}

func health(c *gin.Context) {
	c.String(200, "Ok")
}

func convertTemperature(c *gin.Context) {
    var json temperature

    if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    if json.TemperatureC == "" {
      temperatureF, err := strconv.ParseFloat(json.TemperatureF, 32)
      if err != nil {
        return
      }
      temperatureC :=  (temperatureF - 32) * 0.5556
      c.IndentedJSON(http.StatusOK, temperatureC)	    
    }
    if json.TemperatureF == "" {
      temperatureC, err := strconv.ParseFloat(json.TemperatureC, 32)
      if err != nil {
        return
      }
      temperatureF :=  (temperatureC * 9/5) + 32
      c.IndentedJSON(http.StatusOK, temperatureF)
    }
}

func setup() *gin.Engine {
    router := gin.Default()
    router.GET("/health", health)
    router.POST("/temperature", convertTemperature)

    return router
}

func main() {
	r := setup()
	r.Run(":8080")
}
