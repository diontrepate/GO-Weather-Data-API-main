package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("configurations/openWeatherConfig.env")

	// := used for declaration, it can also assign values
	// This returns a default engine with a logger already attached
	r := gin.Default()

	// The * denotes the variable variable suffixed as a pointer
	r.GET("/weather-data", func(c *gin.Context) {

		// The _ allows us to bypass the other return types
		response, _ := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Kenosha&appid=" + os.Getenv("API_KEY"))

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		c.Data(http.StatusOK, gin.MIMEJSON, responseData)

		fmt.Println("SUCCESS")
	})

	r.Run()
}
