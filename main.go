package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// := used for declaration, it can also assign values
	r := gin.Default()

	// The * denotes the variable variable suffixed as a pointer
	r.GET("/start", func(c *gin.Context) {
		retWeather := Weather{Day: "Monday", Temp: 86}

		actual, err := json.Marshal(retWeather)

		if err != nil {
			log.Fatal(err)
		}

		c.Data(http.StatusOK, gin.MIMEJSON, actual)

		fmt.Println("SUCCESS")
	})

	r.Run()
}

// Made capitolized to allow export ability
type Weather struct {
	Day  string
	Temp int
}
