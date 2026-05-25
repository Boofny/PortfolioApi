package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Boofny/goLive"
	"goliveDocker/rating"
	"github.com/Boofny/goLive/middleware"
)

func main() {
	port, found := os.LookupEnv("PORT")
	if !found {
		log.Fatal("PORT not found in env")
	}
	e := goLive.Launch()

	e.Chain(
		middleware.CORS(),
		middleware.Logger(),
		rating.RateLimit(),
	)

	// Example get req
	e.GET("/ping", func(c *goLive.Context) error {
		return c.SendJSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	//example of a get request with path values
	e.GET("/user/{id}", func(c *goLive.Context) error {
		id := c.Param("id")
		return c.SendSTRING(http.StatusOK, id)
	})
	
	//example of reading json from post request
	e.POST("/posting", func(c *goLive.Context) error {

		type User struct{ 
			Name string `json:"name"`
			Email string `json:"email"`
		}

		var data User
		err := c.ReadJSON(&data)
		if err != nil {
			return c.Error(http.StatusNotFound, "Error in /posting")
		}

		return c.PrettyJSON(http.StatusOK, map[string]any{
			"Name": "Hello " + data.Name,
			"Email": "EMAIL: " + data.Email,
		})

	})

	v1 := e.GroupRoutes("/v1")

	v1.GET("/ping", func(c *goLive.Context) error {
		return c.SendSTRING(http.StatusOK, "pong v1")
	})

	e.StartServer(":"+port)
}
